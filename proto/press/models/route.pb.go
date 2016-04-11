// Code generated by protoc-gen-go.
// source: route.proto
// DO NOT EDIT!

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Route struct {
	Uuid *string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	// Types that are valid to be assigned to Target:
	//	*Route_File
	//	*Route_PageUuid
	Target           isRoute_Target `protobuf_oneof:"target"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isRoute_Target interface {
	isRoute_Target()
}

type Route_File struct {
	File string `protobuf:"bytes,10,opt,name=file,oneof"`
}
type Route_PageUuid struct {
	PageUuid string `protobuf:"bytes,11,opt,name=page_uuid,json=pageUuid,oneof"`
}

func (*Route_File) isRoute_Target()     {}
func (*Route_PageUuid) isRoute_Target() {}

func (m *Route) GetTarget() isRoute_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Route) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *Route) GetFile() string {
	if x, ok := m.GetTarget().(*Route_File); ok {
		return x.File
	}
	return ""
}

func (m *Route) GetPageUuid() string {
	if x, ok := m.GetTarget().(*Route_PageUuid); ok {
		return x.PageUuid
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Route) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Route_OneofMarshaler, _Route_OneofUnmarshaler, _Route_OneofSizer, []interface{}{
		(*Route_File)(nil),
		(*Route_PageUuid)(nil),
	}
}

func _Route_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Route)
	// target
	switch x := m.Target.(type) {
	case *Route_File:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.File)
	case *Route_PageUuid:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.PageUuid)
	case nil:
	default:
		return fmt.Errorf("Route.Target has unexpected type %T", x)
	}
	return nil
}

func _Route_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Route)
	switch tag {
	case 10: // target.file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Target = &Route_File{x}
		return true, err
	case 11: // target.page_uuid
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Target = &Route_PageUuid{x}
		return true, err
	default:
		return false, nil
	}
}

func _Route_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Route)
	// target
	switch x := m.Target.(type) {
	case *Route_File:
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.File)))
		n += len(x.File)
	case *Route_PageUuid:
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.PageUuid)))
		n += len(x.PageUuid)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Route)(nil), "press.models.Route")
}

var fileDescriptor2 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xca, 0x2f, 0x2d,
	0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x28, 0x4a, 0x2d, 0x2e, 0xd6, 0xcb,
	0xcd, 0x4f, 0x49, 0xcd, 0x29, 0x56, 0x8a, 0xe2, 0x62, 0x0d, 0x02, 0x49, 0x0a, 0x09, 0x71, 0xb1,
	0x94, 0x96, 0x66, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x22, 0x5c,
	0x2c, 0x69, 0x99, 0x39, 0xa9, 0x12, 0x5c, 0x20, 0x31, 0x0f, 0x86, 0x20, 0x30, 0x4f, 0x48, 0x96,
	0x8b, 0xb3, 0x20, 0x31, 0x3d, 0x35, 0x1e, 0xac, 0x9c, 0x1b, 0x2a, 0xc5, 0x01, 0x12, 0x0a, 0x05,
	0x8a, 0x38, 0x71, 0x70, 0xb1, 0x95, 0x24, 0x16, 0xa5, 0xa7, 0x96, 0x38, 0xf1, 0x45, 0x41, 0xec,
	0xd2, 0x87, 0xd8, 0x05, 0x08, 0x00, 0x00, 0xff, 0xff, 0x59, 0x8b, 0x17, 0x3c, 0x87, 0x00, 0x00,
	0x00,
}
