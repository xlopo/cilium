// Copyright 2018 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lbmap

import (
	"fmt"

	"github.com/cilium/cilium/pkg/loadbalancer"
	"github.com/cilium/cilium/pkg/lock"
)

type serviceValueMap map[BackendLegacyID]ServiceValue

type bpfBackend struct {
	id       BackendLegacyID
	isHole   bool
	bpfValue ServiceValue
}

type bpfService struct {
	// mutex protects access to all members of bpfService
	mutex lock.RWMutex

	frontendKey ServiceKey

	// backendsByMapIndex is the 1:1 representation of service backends as
	// written into the BPF map. As service backends are scaled up or down,
	// duplicate entries may be required to avoid moving backends to
	// different map index slots. This map represents this and thus may
	// contain duplicate backend entries in different map index slots.
	backendsByMapIndex map[int]*bpfBackend

	// uniqueBackends is a map of all service backends indexed by service
	// backend ID. A backend may be listed multiple times in
	// backendsByMapIndex, it will only be listed once in uniqueBackends.
	uniqueBackends serviceValueMap

	// slaveSlotByBackendLegacyID is a map for getting a position within svc
	// value to any slave which points to a backend identified with
	// the legacy ID.
	slaveSlotByBackendLegacyID map[BackendLegacyID]int

	backendsV2 map[BackendLegacyID]ServiceValue
}

func newBpfService(key ServiceKey) *bpfService {
	return &bpfService{
		frontendKey:                key,
		backendsByMapIndex:         map[int]*bpfBackend{},
		uniqueBackends:             map[BackendLegacyID]ServiceValue{},
		slaveSlotByBackendLegacyID: map[BackendLegacyID]int{},
		backendsV2:                 map[BackendLegacyID]ServiceValue{},
	}
}

func (b *bpfService) addBackend(backend ServiceValue, backendID uint16) int {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	nextSlot := len(b.backendsByMapIndex) + 1
	// TODO(brb) explain hack
	backend.SetCount(int(backendID))
	b.backendsByMapIndex[nextSlot] = &bpfBackend{
		bpfValue: backend,
		id:       backend.BackendLegacyID(),
	}

	b.uniqueBackends[backend.BackendLegacyID()] = backend

	return nextSlot
}

func (b *bpfService) deleteBackend(backend ServiceValue) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// TODO(brb) -> BackendLegacyID
	idToRemove := backend.BackendLegacyID()
	indicesToRemove := []int{}
	duplicateCount := map[BackendLegacyID]int{}

	for index, backend := range b.backendsByMapIndex {
		// create a slice of all backend indices that match the backend
		// ID (ip, port, revnat id)
		if idToRemove == backend.id {
			indicesToRemove = append(indicesToRemove, index)
		} else {
			duplicateCount[backend.id]++
		}
	}

	// select the backend with the most duplicates that is not the backend
	var lowestCount int
	var fillBackendID BackendLegacyID
	for backendID, count := range duplicateCount {
		if lowestCount == 0 || count < lowestCount {
			lowestCount = count
			fillBackendID = backendID
		}
	}

	if fillBackendID == "" {
		// No more entries to fill in, we can remove all backend slots
		b.backendsByMapIndex = map[int]*bpfBackend{}
	} else {
		fillBackend := &bpfBackend{
			id:       fillBackendID,
			isHole:   true,
			bpfValue: b.uniqueBackends[fillBackendID],
		}
		for _, removeIndex := range indicesToRemove {
			b.backendsByMapIndex[removeIndex] = fillBackend
		}
	}

	delete(b.uniqueBackends, idToRemove)
	delete(b.slaveSlotByBackendLegacyID, backend.BackendLegacyID())
}

func (b *bpfService) getSlaveSlot(id BackendLegacyID) (int, bool) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	slot, found := b.slaveSlotByBackendLegacyID[id]
	return slot, found
}

func (b *bpfService) getBackends() []ServiceValue {
	b.mutex.RLock()
	backends := make([]ServiceValue, len(b.backendsByMapIndex))
	dstIndex := 0
	for i := 1; i <= len(b.backendsByMapIndex); i++ {
		if b.backendsByMapIndex[i] == nil {
			log.Errorf("BUG: hole found in backendsByMapIndex: %#v", b.backendsByMapIndex)
			continue
		}

		backends[dstIndex] = b.backendsByMapIndex[i].bpfValue
		dstIndex++
	}
	b.mutex.RUnlock()
	return backends
}

type lbmapCache struct {
	mutex               lock.Mutex
	entries             map[string]*bpfService
	backendRefCount     map[BackendLegacyID]int
	backendIDByLegacyID map[BackendLegacyID]uint16
}

func newLBMapCache() lbmapCache {
	return lbmapCache{
		entries:             map[string]*bpfService{},
		backendRefCount:     map[BackendLegacyID]int{},
		backendIDByLegacyID: map[BackendLegacyID]uint16{},
	}
}

func createBackendsMap(backends []ServiceValue) serviceValueMap {
	m := serviceValueMap{}
	for _, b := range backends {
		m[b.BackendLegacyID()] = b
	}
	return m
}

func (l *lbmapCache) restoreService(svc loadbalancer.LBSVC, v2Exists bool) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	frontendID := svc.FE.String()

	serviceKey, serviceValues, err := LBSVC2ServiceKeynValue(&svc)
	if err != nil {
		return err
	}

	bpfSvc, ok := l.entries[frontendID]
	if !ok {
		bpfSvc = newBpfService(serviceKey)
		l.entries[frontendID] = bpfSvc
	}

	for index, backend := range serviceValues {
		b := &bpfBackend{
			id:       backend.BackendLegacyID(),
			bpfValue: backend,
		}
		if _, ok := bpfSvc.uniqueBackends[backend.BackendLegacyID()]; ok {
			b.isHole = true
		} else {
			bpfSvc.uniqueBackends[backend.BackendLegacyID()] = backend
			bpfSvc.slaveSlotByBackendLegacyID[backend.BackendLegacyID()] = index + 1
		}

		bpfSvc.backendsByMapIndex[index+1] = b
	}

	if v2Exists {
		for legacyID, backend := range bpfSvc.uniqueBackends {
			bpfSvc.backendsV2[legacyID] = backend
			l.addBackendV2Locked(legacyID)
		}
	}

	return nil
}

func (l *lbmapCache) getSlaveSlot(fe ServiceKeyV2, legacyBackendID BackendLegacyID) (int, bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	frontendID := fe.String()
	bpfSvc, found := l.entries[frontendID]
	if !found {
		return 0, false
	}

	pos, found := bpfSvc.slaveSlotByBackendLegacyID[legacyBackendID]
	if !found {
		return 0, false
	}

	return pos, true
}

// assumes that backends doesn't contain frontend service
func (l *lbmapCache) prepareUpdate(fe ServiceKey, backends []ServiceValue) (*bpfService, map[uint16]ServiceValue, []uint16, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	frontendID := fe.String()

	bpfSvc, ok := l.entries[frontendID]
	if !ok {
		bpfSvc = newBpfService(fe)
		l.entries[frontendID] = bpfSvc
	}

	newBackendsMap := createBackendsMap(backends)
	removedBackendIDs := []uint16{}
	addedBackendIDs := map[uint16]ServiceValue{}

	// Step 1: Delete all backends that no longer exist. This will not
	// actually remove the backends but overwrite all slave slots that
	// point to the removed backend with the backend that has the least
	// duplicated slots.
	for legacyID, b := range bpfSvc.uniqueBackends {
		if _, ok := newBackendsMap[legacyID]; !ok {
			bpfSvc.deleteBackend(b)
			delete(bpfSvc.slaveSlotByBackendLegacyID, legacyID)
		}
	}

	for legacyID := range bpfSvc.backendsV2 {
		if _, ok := newBackendsMap[legacyID]; !ok {
			last, err := l.delBackendV2Locked(legacyID)
			if err != nil {
				return nil, nil, nil, err
			}
			if last {
				removedBackendIDs = append(removedBackendIDs, l.backendIDByLegacyID[legacyID])
			}
		}
	}

	// Step 2: Add all backends that don't exist yet.
	for _, b := range backends {
		if _, ok := bpfSvc.uniqueBackends[b.BackendLegacyID()]; !ok {
			legacyID := b.BackendLegacyID()
			backendID := l.backendIDByLegacyID[legacyID]
			pos := bpfSvc.addBackend(b, backendID)
			bpfSvc.slaveSlotByBackendLegacyID[legacyID] = pos
		}
	}

	for _, b := range backends {
		legacyID := b.BackendLegacyID()
		if _, ok := bpfSvc.backendsV2[legacyID]; !ok {
			bpfSvc.backendsV2[legacyID] = b
			first := l.addBackendV2Locked(legacyID)
			if first {
				addedBackendIDs[l.backendIDByLegacyID[legacyID]] = b
			}
		}
	}

	return bpfSvc, addedBackendIDs, removedBackendIDs, nil
}

func (l *lbmapCache) delete(fe ServiceKey) {
	l.mutex.Lock()
	delete(l.entries, fe.String())
	l.mutex.Unlock()
}

// Returns true if new
func (l *lbmapCache) addBackendV2Locked(legacyID BackendLegacyID) bool {
	l.backendRefCount[legacyID]++
	return l.backendRefCount[legacyID] == 1
}

func (l *lbmapCache) existBackendV2(legacyID BackendLegacyID) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.backendRefCount[legacyID] > 0
}

func (l *lbmapCache) delBackendV2Locked(legacyID BackendLegacyID) (bool, error) {
	count, found := l.backendRefCount[legacyID]
	if !found {
		return false, fmt.Errorf("backend %s not found", legacyID)
	}

	if count == 1 {
		delete(l.backendRefCount, legacyID)
		return true, nil
	}

	l.backendRefCount[legacyID]--
	return false, nil
}

func (l *lbmapCache) addBackendIDs(backendIDs map[BackendLegacyID]uint16) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	for legacyID, backendID := range backendIDs {
		l.backendIDByLegacyID[legacyID] = backendID
	}
}

func (l *lbmapCache) missingBackendLegacyIDs(backendLegacyIDs map[BackendLegacyID]struct{}) map[BackendLegacyID]struct{} {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	missing := map[BackendLegacyID]struct{}{}

	for legacyID := range backendLegacyIDs {
		if _, found := l.backendIDByLegacyID[legacyID]; !found {
			missing[legacyID] = struct{}{}
		}
	}

	return missing
}

func (l *lbmapCache) getBackendIDByLegacyID(legacyID BackendLegacyID) uint16 {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	return l.backendIDByLegacyID[legacyID]
}

func (l *lbmapCache) removeServiceV2(svcKey ServiceKeyV2) ([]uint16, int, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	frontendID := svcKey.String()
	bpfSvc, ok := l.entries[frontendID]
	if !ok {
		return nil, 0, fmt.Errorf("Service %s not found", frontendID)
	}

	backendsToRemove := []uint16{}
	count := len(bpfSvc.backendsV2)

	for legacyID := range bpfSvc.backendsV2 {
		last, err := l.delBackendV2Locked(legacyID)
		if err != nil {
			return nil, 0, err
		}
		if last {
			backendsToRemove = append(backendsToRemove, l.backendIDByLegacyID[legacyID])
		}
		delete(l.backendIDByLegacyID, legacyID)
	}

	delete(l.entries, frontendID)

	return backendsToRemove, count, nil
}
