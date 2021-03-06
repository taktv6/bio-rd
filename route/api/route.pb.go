// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bio-routing/bio-rd/route/api/route.proto

package api

import (
	fmt "fmt"
	api "github.com/bio-routing/bio-rd/net/api"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Path_Type int32

const (
	Path_Static Path_Type = 0
	Path_BGP    Path_Type = 1
)

var Path_Type_name = map[int32]string{
	0: "Static",
	1: "BGP",
}

var Path_Type_value = map[string]int32{
	"Static": 0,
	"BGP":    1,
}

func (x Path_Type) String() string {
	return proto.EnumName(Path_Type_name, int32(x))
}

func (Path_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00363871266b6b0e, []int{1, 0}
}

type Route struct {
	Pfx                  *api.Prefix `protobuf:"bytes,1,opt,name=pfx,proto3" json:"pfx,omitempty"`
	Paths                []*Path     `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_00363871266b6b0e, []int{0}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetPfx() *api.Prefix {
	if m != nil {
		return m.Pfx
	}
	return nil
}

func (m *Route) GetPaths() []*Path {
	if m != nil {
		return m.Paths
	}
	return nil
}

type Path struct {
	Type                 Path_Type   `protobuf:"varint,1,opt,name=type,proto3,enum=bio.route.Path_Type" json:"type,omitempty"`
	StaticPath           *StaticPath `protobuf:"bytes,2,opt,name=static_path,json=staticPath,proto3" json:"static_path,omitempty"`
	BgpPath              *BGPPath    `protobuf:"bytes,3,opt,name=bgp_path,json=bgpPath,proto3" json:"bgp_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Path) Reset()         { *m = Path{} }
func (m *Path) String() string { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()    {}
func (*Path) Descriptor() ([]byte, []int) {
	return fileDescriptor_00363871266b6b0e, []int{1}
}

func (m *Path) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Path.Unmarshal(m, b)
}
func (m *Path) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Path.Marshal(b, m, deterministic)
}
func (m *Path) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Path.Merge(m, src)
}
func (m *Path) XXX_Size() int {
	return xxx_messageInfo_Path.Size(m)
}
func (m *Path) XXX_DiscardUnknown() {
	xxx_messageInfo_Path.DiscardUnknown(m)
}

var xxx_messageInfo_Path proto.InternalMessageInfo

func (m *Path) GetType() Path_Type {
	if m != nil {
		return m.Type
	}
	return Path_Static
}

func (m *Path) GetStaticPath() *StaticPath {
	if m != nil {
		return m.StaticPath
	}
	return nil
}

func (m *Path) GetBgpPath() *BGPPath {
	if m != nil {
		return m.BgpPath
	}
	return nil
}

type StaticPath struct {
	NextHop              *api.IP  `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StaticPath) Reset()         { *m = StaticPath{} }
func (m *StaticPath) String() string { return proto.CompactTextString(m) }
func (*StaticPath) ProtoMessage()    {}
func (*StaticPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_00363871266b6b0e, []int{2}
}

func (m *StaticPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StaticPath.Unmarshal(m, b)
}
func (m *StaticPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StaticPath.Marshal(b, m, deterministic)
}
func (m *StaticPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StaticPath.Merge(m, src)
}
func (m *StaticPath) XXX_Size() int {
	return xxx_messageInfo_StaticPath.Size(m)
}
func (m *StaticPath) XXX_DiscardUnknown() {
	xxx_messageInfo_StaticPath.DiscardUnknown(m)
}

var xxx_messageInfo_StaticPath proto.InternalMessageInfo

func (m *StaticPath) GetNextHop() *api.IP {
	if m != nil {
		return m.NextHop
	}
	return nil
}

type BGPPath struct {
	PathIdentifier       uint32                  `protobuf:"varint,1,opt,name=path_identifier,json=pathIdentifier,proto3" json:"path_identifier,omitempty"`
	NextHop              *api.IP                 `protobuf:"bytes,2,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	LocalPref            uint32                  `protobuf:"varint,3,opt,name=local_pref,json=localPref,proto3" json:"local_pref,omitempty"`
	AsPath               []*ASPathSegment        `protobuf:"bytes,4,rep,name=as_path,json=asPath,proto3" json:"as_path,omitempty"`
	Origin               uint32                  `protobuf:"varint,5,opt,name=origin,proto3" json:"origin,omitempty"`
	Med                  uint32                  `protobuf:"varint,6,opt,name=med,proto3" json:"med,omitempty"`
	Ebgp                 bool                    `protobuf:"varint,7,opt,name=ebgp,proto3" json:"ebgp,omitempty"`
	BgpIdentifier        uint32                  `protobuf:"varint,8,opt,name=bgp_identifier,json=bgpIdentifier,proto3" json:"bgp_identifier,omitempty"`
	Source               *api.IP                 `protobuf:"bytes,9,opt,name=source,proto3" json:"source,omitempty"`
	Communities          []uint32                `protobuf:"varint,10,rep,packed,name=communities,proto3" json:"communities,omitempty"`
	LargeCommunities     []*LargeCommunity       `protobuf:"bytes,11,rep,name=large_communities,json=largeCommunities,proto3" json:"large_communities,omitempty"`
	OriginatorId         uint32                  `protobuf:"varint,12,opt,name=originator_id,json=originatorId,proto3" json:"originator_id,omitempty"`
	ClusterList          []uint32                `protobuf:"varint,13,rep,packed,name=cluster_list,json=clusterList,proto3" json:"cluster_list,omitempty"`
	UnknownAttributes    []*UnknownPathAttribute `protobuf:"bytes,14,rep,name=unknown_attributes,json=unknownAttributes,proto3" json:"unknown_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *BGPPath) Reset()         { *m = BGPPath{} }
func (m *BGPPath) String() string { return proto.CompactTextString(m) }
func (*BGPPath) ProtoMessage()    {}
func (*BGPPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_00363871266b6b0e, []int{3}
}

func (m *BGPPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BGPPath.Unmarshal(m, b)
}
func (m *BGPPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BGPPath.Marshal(b, m, deterministic)
}
func (m *BGPPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BGPPath.Merge(m, src)
}
func (m *BGPPath) XXX_Size() int {
	return xxx_messageInfo_BGPPath.Size(m)
}
func (m *BGPPath) XXX_DiscardUnknown() {
	xxx_messageInfo_BGPPath.DiscardUnknown(m)
}

var xxx_messageInfo_BGPPath proto.InternalMessageInfo

func (m *BGPPath) GetPathIdentifier() uint32 {
	if m != nil {
		return m.PathIdentifier
	}
	return 0
}

func (m *BGPPath) GetNextHop() *api.IP {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (m *BGPPath) GetLocalPref() uint32 {
	if m != nil {
		return m.LocalPref
	}
	return 0
}

func (m *BGPPath) GetAsPath() []*ASPathSegment {
	if m != nil {
		return m.AsPath
	}
	return nil
}

func (m *BGPPath) GetOrigin() uint32 {
	if m != nil {
		return m.Origin
	}
	return 0
}

func (m *BGPPath) GetMed() uint32 {
	if m != nil {
		return m.Med
	}
	return 0
}

func (m *BGPPath) GetEbgp() bool {
	if m != nil {
		return m.Ebgp
	}
	return false
}

func (m *BGPPath) GetBgpIdentifier() uint32 {
	if m != nil {
		return m.BgpIdentifier
	}
	return 0
}

func (m *BGPPath) GetSource() *api.IP {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *BGPPath) GetCommunities() []uint32 {
	if m != nil {
		return m.Communities
	}
	return nil
}

func (m *BGPPath) GetLargeCommunities() []*LargeCommunity {
	if m != nil {
		return m.LargeCommunities
	}
	return nil
}

func (m *BGPPath) GetOriginatorId() uint32 {
	if m != nil {
		return m.OriginatorId
	}
	return 0
}

func (m *BGPPath) GetClusterList() []uint32 {
	if m != nil {
		return m.ClusterList
	}
	return nil
}

func (m *BGPPath) GetUnknownAttributes() []*UnknownPathAttribute {
	if m != nil {
		return m.UnknownAttributes
	}
	return nil
}

type ASPathSegment struct {
	AsSequence           bool     `protobuf:"varint,1,opt,name=as_sequence,json=asSequence,proto3" json:"as_sequence,omitempty"`
	Asns                 []uint32 `protobuf:"varint,2,rep,packed,name=asns,proto3" json:"asns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ASPathSegment) Reset()         { *m = ASPathSegment{} }
func (m *ASPathSegment) String() string { return proto.CompactTextString(m) }
func (*ASPathSegment) ProtoMessage()    {}
func (*ASPathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_00363871266b6b0e, []int{4}
}

func (m *ASPathSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ASPathSegment.Unmarshal(m, b)
}
func (m *ASPathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ASPathSegment.Marshal(b, m, deterministic)
}
func (m *ASPathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ASPathSegment.Merge(m, src)
}
func (m *ASPathSegment) XXX_Size() int {
	return xxx_messageInfo_ASPathSegment.Size(m)
}
func (m *ASPathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_ASPathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_ASPathSegment proto.InternalMessageInfo

func (m *ASPathSegment) GetAsSequence() bool {
	if m != nil {
		return m.AsSequence
	}
	return false
}

func (m *ASPathSegment) GetAsns() []uint32 {
	if m != nil {
		return m.Asns
	}
	return nil
}

type LargeCommunity struct {
	GlobalAdministrator  uint32   `protobuf:"varint,1,opt,name=global_administrator,json=globalAdministrator,proto3" json:"global_administrator,omitempty"`
	DataPart1            uint32   `protobuf:"varint,2,opt,name=data_part1,json=dataPart1,proto3" json:"data_part1,omitempty"`
	DataPart2            uint32   `protobuf:"varint,3,opt,name=data_part2,json=dataPart2,proto3" json:"data_part2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LargeCommunity) Reset()         { *m = LargeCommunity{} }
func (m *LargeCommunity) String() string { return proto.CompactTextString(m) }
func (*LargeCommunity) ProtoMessage()    {}
func (*LargeCommunity) Descriptor() ([]byte, []int) {
	return fileDescriptor_00363871266b6b0e, []int{5}
}

func (m *LargeCommunity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LargeCommunity.Unmarshal(m, b)
}
func (m *LargeCommunity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LargeCommunity.Marshal(b, m, deterministic)
}
func (m *LargeCommunity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LargeCommunity.Merge(m, src)
}
func (m *LargeCommunity) XXX_Size() int {
	return xxx_messageInfo_LargeCommunity.Size(m)
}
func (m *LargeCommunity) XXX_DiscardUnknown() {
	xxx_messageInfo_LargeCommunity.DiscardUnknown(m)
}

var xxx_messageInfo_LargeCommunity proto.InternalMessageInfo

func (m *LargeCommunity) GetGlobalAdministrator() uint32 {
	if m != nil {
		return m.GlobalAdministrator
	}
	return 0
}

func (m *LargeCommunity) GetDataPart1() uint32 {
	if m != nil {
		return m.DataPart1
	}
	return 0
}

func (m *LargeCommunity) GetDataPart2() uint32 {
	if m != nil {
		return m.DataPart2
	}
	return 0
}

type UnknownPathAttribute struct {
	Optional             bool     `protobuf:"varint,1,opt,name=optional,proto3" json:"optional,omitempty"`
	Transitive           bool     `protobuf:"varint,2,opt,name=transitive,proto3" json:"transitive,omitempty"`
	Partial              bool     `protobuf:"varint,3,opt,name=partial,proto3" json:"partial,omitempty"`
	TypeCode             uint32   `protobuf:"varint,4,opt,name=type_code,json=typeCode,proto3" json:"type_code,omitempty"`
	Value                []byte   `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnknownPathAttribute) Reset()         { *m = UnknownPathAttribute{} }
func (m *UnknownPathAttribute) String() string { return proto.CompactTextString(m) }
func (*UnknownPathAttribute) ProtoMessage()    {}
func (*UnknownPathAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_00363871266b6b0e, []int{6}
}

func (m *UnknownPathAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnknownPathAttribute.Unmarshal(m, b)
}
func (m *UnknownPathAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnknownPathAttribute.Marshal(b, m, deterministic)
}
func (m *UnknownPathAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnknownPathAttribute.Merge(m, src)
}
func (m *UnknownPathAttribute) XXX_Size() int {
	return xxx_messageInfo_UnknownPathAttribute.Size(m)
}
func (m *UnknownPathAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_UnknownPathAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_UnknownPathAttribute proto.InternalMessageInfo

func (m *UnknownPathAttribute) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

func (m *UnknownPathAttribute) GetTransitive() bool {
	if m != nil {
		return m.Transitive
	}
	return false
}

func (m *UnknownPathAttribute) GetPartial() bool {
	if m != nil {
		return m.Partial
	}
	return false
}

func (m *UnknownPathAttribute) GetTypeCode() uint32 {
	if m != nil {
		return m.TypeCode
	}
	return 0
}

func (m *UnknownPathAttribute) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("bio.route.Path_Type", Path_Type_name, Path_Type_value)
	proto.RegisterType((*Route)(nil), "bio.route.Route")
	proto.RegisterType((*Path)(nil), "bio.route.Path")
	proto.RegisterType((*StaticPath)(nil), "bio.route.StaticPath")
	proto.RegisterType((*BGPPath)(nil), "bio.route.BGPPath")
	proto.RegisterType((*ASPathSegment)(nil), "bio.route.ASPathSegment")
	proto.RegisterType((*LargeCommunity)(nil), "bio.route.LargeCommunity")
	proto.RegisterType((*UnknownPathAttribute)(nil), "bio.route.UnknownPathAttribute")
}

func init() {
	proto.RegisterFile("github.com/bio-routing/bio-rd/route/api/route.proto", fileDescriptor_00363871266b6b0e)
}

var fileDescriptor_00363871266b6b0e = []byte{
	// 730 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xdb, 0x38,
	0x10, 0x5d, 0xc7, 0x5f, 0xf2, 0xd8, 0x72, 0x1c, 0xae, 0x77, 0xa1, 0x4d, 0xb0, 0x8d, 0xa3, 0x20,
	0x8d, 0x7b, 0x88, 0x8d, 0x38, 0x45, 0xef, 0x49, 0x8a, 0xa6, 0x01, 0x82, 0xc2, 0xa5, 0xdb, 0x4b,
	0x2f, 0x02, 0x25, 0xd1, 0x32, 0x51, 0x99, 0x54, 0x45, 0x2a, 0x4d, 0x8e, 0xfd, 0x25, 0xfd, 0x13,
	0x3d, 0xf4, 0xe7, 0x15, 0xa4, 0x64, 0x47, 0xee, 0x17, 0x7a, 0x9b, 0x79, 0x33, 0x6f, 0xe6, 0x3d,
	0x92, 0x12, 0x9c, 0x45, 0x4c, 0x2d, 0x32, 0x7f, 0x14, 0x88, 0xe5, 0xd8, 0x67, 0xe2, 0x24, 0x15,
	0x99, 0x62, 0x3c, 0xca, 0xe3, 0x70, 0xac, 0x53, 0x3a, 0x26, 0x09, 0xcb, 0xa3, 0x51, 0x92, 0x0a,
	0x25, 0x50, 0xcb, 0x67, 0x62, 0x64, 0x80, 0xdd, 0xf1, 0xef, 0xf9, 0x9c, 0x2a, 0xc3, 0xe6, 0x54,
	0xe5, 0x5c, 0xf7, 0x35, 0xd4, 0xb1, 0x66, 0xa2, 0x03, 0xa8, 0x26, 0xf3, 0x3b, 0xa7, 0x32, 0xa8,
	0x0c, 0xdb, 0x93, 0xed, 0x91, 0x1e, 0xa9, 0xbb, 0xa6, 0x29, 0x9d, 0xb3, 0x3b, 0xac, 0x6b, 0xe8,
	0x08, 0xea, 0x09, 0x51, 0x0b, 0xe9, 0x6c, 0x0d, 0xaa, 0xeb, 0xa6, 0x5c, 0xc8, 0x94, 0xa8, 0x05,
	0xce, 0xab, 0xee, 0x97, 0x0a, 0xd4, 0x74, 0x8e, 0x86, 0x50, 0x53, 0xf7, 0x09, 0x35, 0x33, 0xbb,
	0x93, 0xfe, 0x77, 0xed, 0xa3, 0x37, 0xf7, 0x09, 0xc5, 0xa6, 0x03, 0x3d, 0x83, 0xb6, 0x54, 0x44,
	0xb1, 0xc0, 0xd3, 0x23, 0x9c, 0x2d, 0x23, 0xe2, 0x9f, 0x12, 0x61, 0x66, 0xaa, 0x66, 0x0b, 0xc8,
	0x75, 0x8c, 0x4e, 0xc0, 0xf2, 0xa3, 0x24, 0x27, 0x55, 0x0d, 0x09, 0x95, 0x48, 0x17, 0x57, 0x53,
	0xc3, 0x68, 0xfa, 0x51, 0xa2, 0x03, 0x77, 0x0f, 0x6a, 0x7a, 0x29, 0x02, 0x68, 0xe4, 0x03, 0x7b,
	0x7f, 0xa1, 0x26, 0x54, 0x2f, 0xae, 0xa6, 0xbd, 0x8a, 0xfb, 0x14, 0xe0, 0x61, 0x0b, 0x7a, 0x0c,
	0x16, 0xa7, 0x77, 0xca, 0x5b, 0x88, 0xa4, 0x38, 0x93, 0xf6, 0xfa, 0x4c, 0xae, 0xa7, 0xb8, 0xa9,
	0x8b, 0x2f, 0x45, 0xe2, 0x7e, 0xad, 0x41, 0xb3, 0xd8, 0x83, 0x8e, 0x61, 0x5b, 0x2b, 0xf1, 0x58,
	0x48, 0xb9, 0x62, 0x73, 0x46, 0x53, 0x43, 0xb5, 0x71, 0x57, 0xc3, 0xd7, 0x6b, 0x74, 0x63, 0xf8,
	0xd6, 0xaf, 0x87, 0xa3, 0xff, 0x01, 0x62, 0x11, 0x90, 0xd8, 0x4b, 0x52, 0x3a, 0x37, 0x06, 0x6d,
	0xdc, 0x32, 0x88, 0xbe, 0x16, 0x74, 0x0a, 0x4d, 0x22, 0x73, 0xf3, 0x35, 0x73, 0x23, 0x4e, 0xc9,
	0xfc, 0xf9, 0x4c, 0x6b, 0x9a, 0xd1, 0x68, 0x49, 0xb9, 0xc2, 0x0d, 0x22, 0x8d, 0xc4, 0x7f, 0xa1,
	0x21, 0x52, 0x16, 0x31, 0xee, 0xd4, 0xcd, 0xb4, 0x22, 0x43, 0x3d, 0xa8, 0x2e, 0x69, 0xe8, 0x34,
	0x0c, 0xa8, 0x43, 0x84, 0xa0, 0x46, 0xfd, 0x28, 0x71, 0x9a, 0x83, 0xca, 0xd0, 0xc2, 0x26, 0x46,
	0x47, 0xd0, 0xd5, 0xc7, 0x5d, 0xf2, 0x67, 0x19, 0x82, 0xed, 0x47, 0x49, 0xc9, 0xde, 0x21, 0x34,
	0xa4, 0xc8, 0xd2, 0x80, 0x3a, 0xad, 0x1f, 0xcd, 0x15, 0x25, 0x34, 0x80, 0x76, 0x20, 0x96, 0xcb,
	0x8c, 0x33, 0xc5, 0xa8, 0x74, 0x60, 0x50, 0x1d, 0xda, 0xb8, 0x0c, 0xa1, 0x17, 0xb0, 0x13, 0x93,
	0x34, 0xa2, 0x5e, 0xb9, 0xaf, 0x6d, 0x8c, 0xfe, 0x57, 0x32, 0x7a, 0xa3, 0x7b, 0x2e, 0x8b, 0x96,
	0x7b, 0xdc, 0x8b, 0xcb, 0xb9, 0x9e, 0x73, 0x08, 0x76, 0xee, 0x92, 0x28, 0x91, 0x7a, 0x2c, 0x74,
	0x3a, 0x46, 0x74, 0xe7, 0x01, 0xbc, 0x0e, 0xd1, 0x01, 0x74, 0x82, 0x38, 0x93, 0x8a, 0xa6, 0x5e,
	0xcc, 0xa4, 0x72, 0xec, 0x42, 0x4f, 0x8e, 0xdd, 0x30, 0xa9, 0xd0, 0x2b, 0x40, 0x19, 0x7f, 0xcf,
	0xc5, 0x47, 0xee, 0x11, 0xa5, 0x52, 0xe6, 0x67, 0x8a, 0x4a, 0xa7, 0x6b, 0x04, 0xed, 0x97, 0x04,
	0xbd, 0xcd, 0x9b, 0xf4, 0x79, 0x9f, 0xaf, 0xfa, 0xf0, 0x4e, 0x41, 0x5d, 0x23, 0xd2, 0x7d, 0x0e,
	0xf6, 0xc6, 0x25, 0xa1, 0x7d, 0x68, 0x13, 0xe9, 0x49, 0xfa, 0x21, 0xa3, 0x3c, 0xc8, 0x3f, 0x1b,
	0x0b, 0x03, 0x91, 0xb3, 0x02, 0xd1, 0x77, 0x42, 0x24, 0xcf, 0xbf, 0x3f, 0x1b, 0x9b, 0xd8, 0xfd,
	0x54, 0x81, 0xee, 0xe6, 0x11, 0xa0, 0x53, 0xe8, 0x47, 0xb1, 0xf0, 0x49, 0xec, 0x91, 0x70, 0xc9,
	0x38, 0x93, 0x2a, 0xd5, 0x2e, 0x8b, 0xc7, 0xf8, 0x77, 0x5e, 0x3b, 0x2f, 0x97, 0xf4, 0x4b, 0x0b,
	0x89, 0x22, 0x5e, 0x42, 0x52, 0x75, 0x6a, 0xde, 0xa4, 0x8d, 0x5b, 0x1a, 0x99, 0x6a, 0x60, 0xa3,
	0x3c, 0x59, 0x3d, 0xc4, 0x55, 0x79, 0xe2, 0x7e, 0xae, 0x40, 0xff, 0x67, 0xae, 0xd1, 0x2e, 0x58,
	0x22, 0x51, 0x4c, 0x70, 0x12, 0x17, 0x76, 0xd6, 0x39, 0x7a, 0x04, 0xa0, 0x52, 0xc2, 0x25, 0x53,
	0xec, 0x96, 0x9a, 0x95, 0x16, 0x2e, 0x21, 0xc8, 0x81, 0xa6, 0x5e, 0xc7, 0x48, 0x6c, 0x16, 0x5a,
	0x78, 0x95, 0xa2, 0x3d, 0x68, 0xe9, 0xbf, 0x86, 0x17, 0x88, 0x90, 0x3a, 0x35, 0x23, 0xc6, 0xd2,
	0xc0, 0xa5, 0x08, 0x29, 0xea, 0x43, 0xfd, 0x96, 0xc4, 0x19, 0x35, 0x0f, 0xbc, 0x83, 0xf3, 0xe4,
	0xe2, 0xc9, 0xbb, 0xe3, 0x3f, 0xfc, 0xb3, 0xfa, 0x0d, 0xf3, 0x63, 0x3c, 0xfb, 0x16, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0xbe, 0x17, 0xdc, 0x8b, 0x05, 0x00, 0x00,
}
