// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/listener/listener.proto

package listener

import (
	fmt "fmt"
	auth "github.com/cilium/proxy/go/envoy/api/v2/auth"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/lyft/protoc-gen-validate/validate"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Filter struct {
	// The name of the filter to instantiate. The name must match a supported
	// filter. The built-in filters are:
	//
	// [#comment:TODO(mattklein123): Auto generate the following list]
	// * :ref:`envoy.client_ssl_auth<config_network_filters_client_ssl_auth>`
	// * :ref:`envoy.echo <config_network_filters_echo>`
	// * :ref:`envoy.http_connection_manager <config_http_conn_man>`
	// * :ref:`envoy.mongo_proxy <config_network_filters_mongo_proxy>`
	// * :ref:`envoy.ratelimit <config_network_filters_rate_limit>`
	// * :ref:`envoy.redis_proxy <config_network_filters_redis_proxy>`
	// * :ref:`envoy.tcp_proxy <config_network_filters_tcp_proxy>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being
	// instantiated. See the supported filters for further documentation.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*Filter_Config
	//	*Filter_TypedConfig
	ConfigType isFilter_ConfigType `protobuf_oneof:"config_type"`
	// [#not-implemented-hide:]
	DeprecatedV1         *Filter_DeprecatedV1 `protobuf:"bytes,3,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isFilter_ConfigType interface {
	isFilter_ConfigType()
}

type Filter_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type Filter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,4,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Filter_Config) isFilter_ConfigType() {}

func (*Filter_TypedConfig) isFilter_ConfigType() {}

func (m *Filter) GetConfigType() isFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *Filter) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*Filter_Config); ok {
		return x.Config
	}
	return nil
}

func (m *Filter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Filter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// Deprecated: Do not use.
func (m *Filter) GetDeprecatedV1() *Filter_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Filter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Filter_OneofMarshaler, _Filter_OneofUnmarshaler, _Filter_OneofSizer, []interface{}{
		(*Filter_Config)(nil),
		(*Filter_TypedConfig)(nil),
	}
}

func _Filter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Filter)
	// config_type
	switch x := m.ConfigType.(type) {
	case *Filter_Config:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *Filter_TypedConfig:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TypedConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Filter.ConfigType has unexpected type %T", x)
	}
	return nil
}

func _Filter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Filter)
	switch tag {
	case 2: // config_type.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(_struct.Struct)
		err := b.DecodeMessage(msg)
		m.ConfigType = &Filter_Config{msg}
		return true, err
	case 4: // config_type.typed_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(any.Any)
		err := b.DecodeMessage(msg)
		m.ConfigType = &Filter_TypedConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Filter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Filter)
	// config_type
	switch x := m.ConfigType.(type) {
	case *Filter_Config:
		s := proto.Size(x.Config)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Filter_TypedConfig:
		s := proto.Size(x.TypedConfig)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// [#not-implemented-hide:]
type Filter_DeprecatedV1 struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter_DeprecatedV1) Reset()         { *m = Filter_DeprecatedV1{} }
func (m *Filter_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*Filter_DeprecatedV1) ProtoMessage()    {}
func (*Filter_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{0, 0}
}

func (m *Filter_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter_DeprecatedV1.Unmarshal(m, b)
}
func (m *Filter_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *Filter_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter_DeprecatedV1.Merge(m, src)
}
func (m *Filter_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_Filter_DeprecatedV1.Size(m)
}
func (m *Filter_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_Filter_DeprecatedV1 proto.InternalMessageInfo

func (m *Filter_DeprecatedV1) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// Specifies the match criteria for selecting a specific filter chain for a
// listener.
//
// In order for a filter chain to be selected, *ALL* of its criteria must be
// fulfilled by the incoming connection, properties of which are set by the
// networking stack and/or listener filters.
//
// The following order applies:
//
// 1. Destination port.
// 2. Destination IP address.
// 3. Server name (e.g. SNI for TLS protocol),
// 4. Transport protocol.
// 5. Application protocols (e.g. ALPN for TLS protocol).
//
// For criteria that allow ranges or wildcards, the most specific value in any
// of the configured filter chains that matches the incoming connection is going
// to be used (e.g. for SNI ``www.example.com`` the most specific match would be
// ``www.example.com``, then ``*.example.com``, then ``*.com``, then any filter
// chain without ``server_names`` requirements).
//
// [#comment: Implemented rules are kept in the preference order, with deprecated fields
// listed at the end, because that's how we want to list them in the docs.
//
// [#comment:TODO(PiotrSikora): Add support for configurable precedence of the rules]
type FilterChainMatch struct {
	// Optional destination port to consider when use_original_dst is set on the
	// listener in determining a filter chain match.
	DestinationPort *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	// If non-empty, an IP address and prefix length to match addresses when the
	// listener is bound to 0.0.0.0/:: or when use_original_dst is specified.
	PrefixRanges []*core.CidrRange `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges,proto3" json:"prefix_ranges,omitempty"`
	// If non-empty, an IP address and suffix length to match addresses when the
	// listener is bound to 0.0.0.0/:: or when use_original_dst is specified.
	// [#not-implemented-hide:]
	AddressSuffix string `protobuf:"bytes,4,opt,name=address_suffix,json=addressSuffix,proto3" json:"address_suffix,omitempty"`
	// [#not-implemented-hide:]
	SuffixLen *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=suffix_len,json=suffixLen,proto3" json:"suffix_len,omitempty"`
	// The criteria is satisfied if the source IP address of the downstream
	// connection is contained in at least one of the specified subnets. If the
	// parameter is not specified or the list is empty, the source IP address is
	// ignored.
	// [#not-implemented-hide:]
	SourcePrefixRanges []*core.CidrRange `protobuf:"bytes,6,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	// The criteria is satisfied if the source port of the downstream connection
	// is contained in at least one of the specified ports. If the parameter is
	// not specified, the source port is ignored.
	// [#not-implemented-hide:]
	SourcePorts []*wrappers.UInt32Value `protobuf:"bytes,7,rep,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	// If non-empty, a list of server names (e.g. SNI for TLS protocol) to consider when determining
	// a filter chain match. Those values will be compared against the server names of a new
	// connection, when detected by one of the listener filters.
	//
	// The server name will be matched against all wildcard domains, i.e. ``www.example.com``
	// will be first matched against ``www.example.com``, then ``*.example.com``, then ``*.com``.
	//
	// Note that partial wildcards are not supported, and values like ``*w.example.com`` are invalid.
	//
	// .. attention::
	//
	//   See the :ref:`FAQ entry <faq_how_to_setup_sni>` on how to configure SNI for more
	//   information.
	ServerNames []string `protobuf:"bytes,11,rep,name=server_names,json=serverNames,proto3" json:"server_names,omitempty"`
	// If non-empty, a transport protocol to consider when determining a filter chain match.
	// This value will be compared against the transport protocol of a new connection, when
	// it's detected by one of the listener filters.
	//
	// Suggested values include:
	//
	// * ``raw_buffer`` - default, used when no transport protocol is detected,
	// * ``tls`` - set by :ref:`envoy.listener.tls_inspector <config_listener_filters_tls_inspector>`
	//   when TLS protocol is detected.
	TransportProtocol string `protobuf:"bytes,9,opt,name=transport_protocol,json=transportProtocol,proto3" json:"transport_protocol,omitempty"`
	// If non-empty, a list of application protocols (e.g. ALPN for TLS protocol) to consider when
	// determining a filter chain match. Those values will be compared against the application
	// protocols of a new connection, when detected by one of the listener filters.
	//
	// Suggested values include:
	//
	// * ``http/1.1`` - set by :ref:`envoy.listener.tls_inspector
	//   <config_listener_filters_tls_inspector>`,
	// * ``h2`` - set by :ref:`envoy.listener.tls_inspector <config_listener_filters_tls_inspector>`
	//
	// .. attention::
	//
	//   Currently, only :ref:`TLS Inspector <config_listener_filters_tls_inspector>` provides
	//   application protocol detection based on the requested
	//   `ALPN <https://en.wikipedia.org/wiki/Application-Layer_Protocol_Negotiation>`_ values.
	//
	//   However, the use of ALPN is pretty much limited to the HTTP/2 traffic on the Internet,
	//   and matching on values other than ``h2`` is going to lead to a lot of false negatives,
	//   unless all connecting clients are known to use ALPN.
	ApplicationProtocols []string `protobuf:"bytes,10,rep,name=application_protocols,json=applicationProtocols,proto3" json:"application_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterChainMatch) Reset()         { *m = FilterChainMatch{} }
func (m *FilterChainMatch) String() string { return proto.CompactTextString(m) }
func (*FilterChainMatch) ProtoMessage()    {}
func (*FilterChainMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{1}
}

func (m *FilterChainMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChainMatch.Unmarshal(m, b)
}
func (m *FilterChainMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChainMatch.Marshal(b, m, deterministic)
}
func (m *FilterChainMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChainMatch.Merge(m, src)
}
func (m *FilterChainMatch) XXX_Size() int {
	return xxx_messageInfo_FilterChainMatch.Size(m)
}
func (m *FilterChainMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChainMatch.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChainMatch proto.InternalMessageInfo

func (m *FilterChainMatch) GetDestinationPort() *wrappers.UInt32Value {
	if m != nil {
		return m.DestinationPort
	}
	return nil
}

func (m *FilterChainMatch) GetPrefixRanges() []*core.CidrRange {
	if m != nil {
		return m.PrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetAddressSuffix() string {
	if m != nil {
		return m.AddressSuffix
	}
	return ""
}

func (m *FilterChainMatch) GetSuffixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.SuffixLen
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePrefixRanges() []*core.CidrRange {
	if m != nil {
		return m.SourcePrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePorts() []*wrappers.UInt32Value {
	if m != nil {
		return m.SourcePorts
	}
	return nil
}

func (m *FilterChainMatch) GetServerNames() []string {
	if m != nil {
		return m.ServerNames
	}
	return nil
}

func (m *FilterChainMatch) GetTransportProtocol() string {
	if m != nil {
		return m.TransportProtocol
	}
	return ""
}

func (m *FilterChainMatch) GetApplicationProtocols() []string {
	if m != nil {
		return m.ApplicationProtocols
	}
	return nil
}

// A filter chain wraps a set of match criteria, an option TLS context, a set of filters, and
// various other parameters.
type FilterChain struct {
	// The criteria to use when matching a connection to this filter chain.
	FilterChainMatch *FilterChainMatch `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch,proto3" json:"filter_chain_match,omitempty"`
	// The TLS context for this filter chain.
	TlsContext *auth.DownstreamTlsContext `protobuf:"bytes,2,opt,name=tls_context,json=tlsContext,proto3" json:"tls_context,omitempty"`
	// A list of individual network filters that make up the filter chain for
	// connections established with the listener. Order matters as the filters are
	// processed sequentially as connection events happen. Note: If the filter
	// list is empty, the connection will close by default.
	Filters []*Filter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	// Whether the listener should expect a PROXY protocol V1 header on new
	// connections. If this option is enabled, the listener will assume that that
	// remote address of the connection is the one specified in the header. Some
	// load balancers including the AWS ELB support this option. If the option is
	// absent or set to false, Envoy will use the physical peer address of the
	// connection as the remote address.
	UseProxyProto *wrappers.BoolValue `protobuf:"bytes,4,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	// [#not-implemented-hide:] filter chain metadata.
	Metadata *core.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// See :ref:`base.TransportSocket<envoy_api_msg_core.TransportSocket>` description.
	TransportSocket      *core.TransportSocket `protobuf:"bytes,6,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FilterChain) Reset()         { *m = FilterChain{} }
func (m *FilterChain) String() string { return proto.CompactTextString(m) }
func (*FilterChain) ProtoMessage()    {}
func (*FilterChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{2}
}

func (m *FilterChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChain.Unmarshal(m, b)
}
func (m *FilterChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChain.Marshal(b, m, deterministic)
}
func (m *FilterChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChain.Merge(m, src)
}
func (m *FilterChain) XXX_Size() int {
	return xxx_messageInfo_FilterChain.Size(m)
}
func (m *FilterChain) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChain.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChain proto.InternalMessageInfo

func (m *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if m != nil {
		return m.FilterChainMatch
	}
	return nil
}

func (m *FilterChain) GetTlsContext() *auth.DownstreamTlsContext {
	if m != nil {
		return m.TlsContext
	}
	return nil
}

func (m *FilterChain) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *FilterChain) GetUseProxyProto() *wrappers.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *FilterChain) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FilterChain) GetTransportSocket() *core.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

type ListenerFilter struct {
	// The name of the filter to instantiate. The name must match a supported
	// filter. The built-in filters are:
	//
	// [#comment:TODO(mattklein123): Auto generate the following list]
	// * :ref:`envoy.listener.original_dst <config_listener_filters_original_dst>`
	// * :ref:`envoy.listener.tls_inspector <config_listener_filters_tls_inspector>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being instantiated.
	// See the supported filters for further documentation.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*ListenerFilter_Config
	//	*ListenerFilter_TypedConfig
	ConfigType           isListenerFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ListenerFilter) Reset()         { *m = ListenerFilter{} }
func (m *ListenerFilter) String() string { return proto.CompactTextString(m) }
func (*ListenerFilter) ProtoMessage()    {}
func (*ListenerFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{3}
}

func (m *ListenerFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilter.Unmarshal(m, b)
}
func (m *ListenerFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilter.Marshal(b, m, deterministic)
}
func (m *ListenerFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilter.Merge(m, src)
}
func (m *ListenerFilter) XXX_Size() int {
	return xxx_messageInfo_ListenerFilter.Size(m)
}
func (m *ListenerFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilter proto.InternalMessageInfo

func (m *ListenerFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isListenerFilter_ConfigType interface {
	isListenerFilter_ConfigType()
}

type ListenerFilter_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type ListenerFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ListenerFilter_Config) isListenerFilter_ConfigType() {}

func (*ListenerFilter_TypedConfig) isListenerFilter_ConfigType() {}

func (m *ListenerFilter) GetConfigType() isListenerFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ListenerFilter) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*ListenerFilter_Config); ok {
		return x.Config
	}
	return nil
}

func (m *ListenerFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ListenerFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ListenerFilter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ListenerFilter_OneofMarshaler, _ListenerFilter_OneofUnmarshaler, _ListenerFilter_OneofSizer, []interface{}{
		(*ListenerFilter_Config)(nil),
		(*ListenerFilter_TypedConfig)(nil),
	}
}

func _ListenerFilter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ListenerFilter)
	// config_type
	switch x := m.ConfigType.(type) {
	case *ListenerFilter_Config:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *ListenerFilter_TypedConfig:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TypedConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ListenerFilter.ConfigType has unexpected type %T", x)
	}
	return nil
}

func _ListenerFilter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ListenerFilter)
	switch tag {
	case 2: // config_type.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(_struct.Struct)
		err := b.DecodeMessage(msg)
		m.ConfigType = &ListenerFilter_Config{msg}
		return true, err
	case 3: // config_type.typed_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(any.Any)
		err := b.DecodeMessage(msg)
		m.ConfigType = &ListenerFilter_TypedConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ListenerFilter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ListenerFilter)
	// config_type
	switch x := m.ConfigType.(type) {
	case *ListenerFilter_Config:
		s := proto.Size(x.Config)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ListenerFilter_TypedConfig:
		s := proto.Size(x.TypedConfig)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Filter)(nil), "envoy.api.v2.listener.Filter")
	proto.RegisterType((*Filter_DeprecatedV1)(nil), "envoy.api.v2.listener.Filter.DeprecatedV1")
	proto.RegisterType((*FilterChainMatch)(nil), "envoy.api.v2.listener.FilterChainMatch")
	proto.RegisterType((*FilterChain)(nil), "envoy.api.v2.listener.FilterChain")
	proto.RegisterType((*ListenerFilter)(nil), "envoy.api.v2.listener.ListenerFilter")
}

func init() {
	proto.RegisterFile("envoy/api/v2/listener/listener.proto", fileDescriptor_0ced541f18749edd)
}

var fileDescriptor_0ced541f18749edd = []byte{
	// 820 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4f, 0x6f, 0x23, 0x35,
	0x1c, 0xed, 0x24, 0x21, 0x9b, 0x78, 0x92, 0x6d, 0xb0, 0xba, 0xea, 0x50, 0xf6, 0x4f, 0x88, 0x40,
	0x44, 0x2b, 0x31, 0xa3, 0xa6, 0x07, 0x84, 0x10, 0x42, 0x9d, 0x2e, 0xa8, 0xa0, 0x6d, 0x55, 0x39,
	0xbb, 0x3d, 0x70, 0x19, 0xb9, 0x33, 0x4e, 0x6a, 0x31, 0xb1, 0x47, 0xb6, 0x27, 0xdb, 0x7c, 0x18,
	0xee, 0x08, 0xf1, 0x05, 0x40, 0x1c, 0x38, 0xf2, 0x29, 0x38, 0x70, 0xe3, 0x53, 0x2c, 0xf2, 0x9f,
	0xc9, 0x26, 0x69, 0x54, 0xf6, 0xb4, 0x37, 0xfb, 0xf7, 0x7b, 0xef, 0xcd, 0xfb, 0xd9, 0x6f, 0x0c,
	0x3e, 0x26, 0x6c, 0xce, 0x17, 0x11, 0x2e, 0x68, 0x34, 0x1f, 0x45, 0x39, 0x95, 0x8a, 0x30, 0x22,
	0x96, 0x8b, 0xb0, 0x10, 0x5c, 0x71, 0xf8, 0xc0, 0xa0, 0x42, 0x5c, 0xd0, 0x70, 0x3e, 0x0a, 0xab,
	0xe6, 0xc1, 0x93, 0x35, 0x72, 0xca, 0x05, 0x89, 0x70, 0x96, 0x09, 0x22, 0xa5, 0xe5, 0x1d, 0x3c,
	0x5c, 0x03, 0xe0, 0x52, 0x5d, 0x47, 0x29, 0x11, 0x6a, 0x6b, 0xd7, 0xd0, 0xaf, 0xb0, 0x24, 0xae,
	0xfb, 0xc1, 0x94, 0xf3, 0x69, 0x4e, 0x22, 0xb3, 0xbb, 0x2a, 0x27, 0x11, 0x66, 0x8b, 0x8a, 0xb8,
	0xd9, 0x92, 0x4a, 0x94, 0x69, 0x25, 0xfb, 0x78, 0xb3, 0xfb, 0x4a, 0xe0, 0xa2, 0x20, 0xa2, 0x32,
	0xb5, 0x3f, 0xc7, 0x39, 0xcd, 0xb0, 0x22, 0x51, 0xb5, 0x70, 0x8d, 0xbd, 0x29, 0x9f, 0x72, 0xb3,
	0x8c, 0xf4, 0xca, 0x56, 0x07, 0x3f, 0xd5, 0x40, 0xf3, 0x5b, 0x9a, 0x2b, 0x22, 0xe0, 0x23, 0xd0,
	0x60, 0x78, 0x46, 0x02, 0xaf, 0xef, 0x0d, 0xdb, 0x71, 0xfb, 0xf7, 0x7f, 0xff, 0xac, 0x37, 0x44,
	0xad, 0xef, 0x21, 0x53, 0x86, 0x87, 0xa0, 0x99, 0x72, 0x36, 0xa1, 0xd3, 0xa0, 0xd6, 0xf7, 0x86,
	0xfe, 0x68, 0x3f, 0xb4, 0x4e, 0xc2, 0xca, 0x49, 0x38, 0x36, 0x3e, 0x4f, 0x77, 0x90, 0x03, 0xc2,
	0x2f, 0x40, 0x47, 0x2d, 0x0a, 0x92, 0x25, 0x8e, 0xd8, 0x30, 0xc4, 0xbd, 0x5b, 0xc4, 0x63, 0xb6,
	0x38, 0xdd, 0x41, 0xbe, 0xc1, 0x9e, 0x58, 0xea, 0x18, 0x74, 0x33, 0x52, 0x08, 0x92, 0x62, 0x45,
	0xb2, 0x64, 0x7e, 0x18, 0xd4, 0x0d, 0xf7, 0x69, 0xb8, 0xf5, 0xae, 0x42, 0x3b, 0x42, 0xf8, 0x6c,
	0x49, 0xb9, 0x3c, 0x8c, 0x6b, 0x81, 0x87, 0x3a, 0xd9, 0x4a, 0xe5, 0x60, 0x00, 0x3a, 0xab, 0x08,
	0x08, 0x41, 0x43, 0x7f, 0xd3, 0x4e, 0x8c, 0xcc, 0x3a, 0xee, 0x02, 0xdf, 0xba, 0x4d, 0xf4, 0x76,
	0xf0, 0x47, 0x03, 0xf4, 0xac, 0xf8, 0xc9, 0x35, 0xa6, 0xec, 0x0c, 0xab, 0xf4, 0x1a, 0x5e, 0x82,
	0x5e, 0x46, 0xa4, 0xa2, 0x0c, 0x2b, 0xca, 0x59, 0x52, 0x70, 0xa1, 0x82, 0x96, 0xf1, 0xf7, 0xf0,
	0xd6, 0x6c, 0x2f, 0xbf, 0x63, 0xea, 0x68, 0x74, 0x89, 0xf3, 0x92, 0xc4, 0x5d, 0x7d, 0xa6, 0xad,
	0xa7, 0xcd, 0xe0, 0xf5, 0xeb, 0xfa, 0xd0, 0x43, 0xbb, 0x2b, 0x22, 0x17, 0x5c, 0x28, 0x78, 0x0c,
	0xba, 0x85, 0x20, 0x13, 0x7a, 0x93, 0x08, 0xcc, 0xa6, 0x44, 0x06, 0xf5, 0x7e, 0xdd, 0x88, 0xae,
	0x0d, 0xad, 0xa3, 0x14, 0x9e, 0xd0, 0x4c, 0x20, 0x0d, 0x42, 0x1d, 0x4b, 0x31, 0x1b, 0x09, 0x3f,
	0x01, 0xf7, 0x5d, 0x48, 0x13, 0x59, 0x4e, 0x26, 0xf4, 0xc6, 0x1c, 0x7a, 0x1b, 0x75, 0x5d, 0x75,
	0x6c, 0x8a, 0xf0, 0x4b, 0x00, 0x6c, 0x3b, 0xc9, 0x09, 0x0b, 0xde, 0xfb, 0x7f, 0xef, 0xa8, 0x6d,
	0xf1, 0xcf, 0x09, 0x83, 0xe7, 0x60, 0x4f, 0xf2, 0x52, 0xa4, 0x24, 0x59, 0x77, 0xdb, 0x7c, 0x0b,
	0xb7, 0xd0, 0x32, 0x2f, 0x56, 0x3d, 0x7f, 0x0d, 0x3a, 0x95, 0x1e, 0x17, 0x4a, 0x06, 0xf7, 0x9c,
	0xce, 0x5d, 0x76, 0x7c, 0xa7, 0xa3, 0x09, 0xf0, 0x23, 0xd0, 0x91, 0x44, 0xcc, 0x89, 0x48, 0x74,
	0x52, 0x65, 0xe0, 0xf7, 0xeb, 0xc3, 0x36, 0xf2, 0x6d, 0xed, 0x5c, 0x97, 0xe0, 0x67, 0x00, 0x2a,
	0x81, 0x99, 0xd4, 0x5f, 0x48, 0x8c, 0x62, 0xca, 0xf3, 0xa0, 0x6d, 0xce, 0xe6, 0xfd, 0x65, 0xe7,
	0xc2, 0x35, 0xe0, 0x11, 0x78, 0x80, 0x8b, 0x22, 0xa7, 0xa9, 0xbb, 0x61, 0x57, 0x97, 0x01, 0x30,
	0xd2, 0x7b, 0x2b, 0xcd, 0x8a, 0x23, 0xbf, 0x6f, 0xb4, 0xbc, 0x5e, 0x0d, 0xf9, 0x92, 0xd1, 0x24,
	0xe3, 0x33, 0x4c, 0x99, 0x1c, 0xfc, 0x56, 0x07, 0xfe, 0x4a, 0x7c, 0xe0, 0x4b, 0x00, 0x27, 0x66,
	0x9b, 0xa4, 0x7a, 0x9f, 0xcc, 0x74, 0x9e, 0x4c, 0xfe, 0xfc, 0xd1, 0xa7, 0x77, 0x66, 0xfb, 0x4d,
	0xfc, 0x50, 0x6f, 0xb2, 0x19, 0xc8, 0x53, 0xe0, 0xab, 0x5c, 0xea, 0xdf, 0x4c, 0x91, 0x1b, 0xe5,
	0x7e, 0xd0, 0x0d, 0x3d, 0xfd, 0x3e, 0x85, 0xcf, 0xf8, 0x2b, 0x26, 0x95, 0x20, 0x78, 0xf6, 0x22,
	0x97, 0x27, 0x16, 0x8e, 0x80, 0x5a, 0xae, 0xe1, 0x57, 0xe0, 0x9e, 0x55, 0xaf, 0xc2, 0xf7, 0xe8,
	0x4e, 0x57, 0x71, 0xe3, 0xaf, 0xbf, 0x9f, 0xec, 0xa0, 0x8a, 0x03, 0x63, 0xb0, 0x5b, 0x4a, 0x9d,
	0x0b, 0x7e, 0xb3, 0xb0, 0xa7, 0xe6, 0x7e, 0xfa, 0x83, 0x5b, 0xb7, 0x19, 0x73, 0x9e, 0xdb, 0xbb,
	0xec, 0x96, 0x92, 0x5c, 0x68, 0x86, 0x39, 0x4a, 0xf8, 0x39, 0x68, 0xcd, 0x88, 0xc2, 0x19, 0x56,
	0xd8, 0x25, 0xf3, 0xc3, 0x2d, 0x91, 0x3a, 0x73, 0x10, 0xb4, 0x04, 0xc3, 0x33, 0xd0, 0x7b, 0x73,
	0xc7, 0x92, 0xa7, 0x3f, 0x12, 0x15, 0x34, 0x8d, 0xc0, 0x60, 0x8b, 0xc0, 0x8b, 0x0a, 0x3a, 0x36,
	0x48, 0xb4, 0xab, 0xd6, 0x0b, 0x83, 0x5f, 0x3d, 0x70, 0xff, 0xb9, 0x1b, 0xf7, 0x9d, 0x3d, 0x91,
	0xf5, 0xb7, 0x7e, 0x22, 0x37, 0x5e, 0xaa, 0x78, 0xf8, 0xf3, 0x3f, 0x8f, 0xbd, 0x1f, 0x5a, 0xd5,
	0x05, 0xfd, 0x52, 0xdb, 0xff, 0xc6, 0xcc, 0x7c, 0x5c, 0xd0, 0xf0, 0x72, 0x14, 0x56, 0x93, 0x9c,
	0x8f, 0xaf, 0x9a, 0x46, 0xf5, 0xe8, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x7f, 0x8a, 0x64,
	0x1e, 0x07, 0x00, 0x00,
}