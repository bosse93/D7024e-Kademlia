// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	data.proto

It has these top-level messages:
	RequestPing
	RequestContact
	RequestData
	RequestStore
	Reply
	ReplyContact
	ReplyData
	WrapperMessage
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestPing struct {
	Id string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
}

func (m *RequestPing) Reset()                    { *m = RequestPing{} }
func (m *RequestPing) String() string            { return proto.CompactTextString(m) }
func (*RequestPing) ProtoMessage()               {}
func (*RequestPing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RequestPing) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RequestContact struct {
	Id     string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Target string `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
}

func (m *RequestContact) Reset()                    { *m = RequestContact{} }
func (m *RequestContact) String() string            { return proto.CompactTextString(m) }
func (*RequestContact) ProtoMessage()               {}
func (*RequestContact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RequestContact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RequestContact) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type RequestData struct {
	Id  string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=Key" json:"Key,omitempty"`
}

func (m *RequestData) Reset()                    { *m = RequestData{} }
func (m *RequestData) String() string            { return proto.CompactTextString(m) }
func (*RequestData) ProtoMessage()               {}
func (*RequestData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RequestData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RequestData) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type RequestStore struct {
	Id   string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=Key" json:"Key,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=Data" json:"Data,omitempty"`
}

func (m *RequestStore) Reset()                    { *m = RequestStore{} }
func (m *RequestStore) String() string            { return proto.CompactTextString(m) }
func (*RequestStore) ProtoMessage()               {}
func (*RequestStore) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestStore) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RequestStore) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestStore) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Reply struct {
	Id   string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Reply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Reply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ReplyContact struct {
	Id       string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Contacts []*ReplyContact_Contact `protobuf:"bytes,2,rep,name=Contacts" json:"Contacts,omitempty"`
}

func (m *ReplyContact) Reset()                    { *m = ReplyContact{} }
func (m *ReplyContact) String() string            { return proto.CompactTextString(m) }
func (*ReplyContact) ProtoMessage()               {}
func (*ReplyContact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReplyContact) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReplyContact) GetContacts() []*ReplyContact_Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type ReplyContact_Contact struct {
	ID       string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=Address" json:"Address,omitempty"`
	Distance string `protobuf:"bytes,3,opt,name=Distance" json:"Distance,omitempty"`
}

func (m *ReplyContact_Contact) Reset()                    { *m = ReplyContact_Contact{} }
func (m *ReplyContact_Contact) String() string            { return proto.CompactTextString(m) }
func (*ReplyContact_Contact) ProtoMessage()               {}
func (*ReplyContact_Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *ReplyContact_Contact) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ReplyContact_Contact) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ReplyContact_Contact) GetDistance() string {
	if m != nil {
		return m.Distance
	}
	return ""
}

type ReplyData struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ReturnType string `protobuf:"bytes,2,opt,name=returnType" json:"returnType,omitempty"`
	// Types that are valid to be assigned to Msg:
	//	*ReplyData_ReplyData
	//	*ReplyData_ReplyContact
	Msg isReplyData_Msg `protobuf_oneof:"msg"`
}

func (m *ReplyData) Reset()                    { *m = ReplyData{} }
func (m *ReplyData) String() string            { return proto.CompactTextString(m) }
func (*ReplyData) ProtoMessage()               {}
func (*ReplyData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isReplyData_Msg interface {
	isReplyData_Msg()
}

type ReplyData_ReplyData struct {
	ReplyData *Reply `protobuf:"bytes,3,opt,name=replyData,oneof"`
}
type ReplyData_ReplyContact struct {
	ReplyContact *ReplyContact `protobuf:"bytes,4,opt,name=replyContact,oneof"`
}

func (*ReplyData_ReplyData) isReplyData_Msg()    {}
func (*ReplyData_ReplyContact) isReplyData_Msg() {}

func (m *ReplyData) GetMsg() isReplyData_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ReplyData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReplyData) GetReturnType() string {
	if m != nil {
		return m.ReturnType
	}
	return ""
}

func (m *ReplyData) GetReplyData() *Reply {
	if x, ok := m.GetMsg().(*ReplyData_ReplyData); ok {
		return x.ReplyData
	}
	return nil
}

func (m *ReplyData) GetReplyContact() *ReplyContact {
	if x, ok := m.GetMsg().(*ReplyData_ReplyContact); ok {
		return x.ReplyContact
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ReplyData) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ReplyData_OneofMarshaler, _ReplyData_OneofUnmarshaler, _ReplyData_OneofSizer, []interface{}{
		(*ReplyData_ReplyData)(nil),
		(*ReplyData_ReplyContact)(nil),
	}
}

func _ReplyData_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ReplyData)
	// msg
	switch x := m.Msg.(type) {
	case *ReplyData_ReplyData:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReplyData); err != nil {
			return err
		}
	case *ReplyData_ReplyContact:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReplyContact); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ReplyData.Msg has unexpected type %T", x)
	}
	return nil
}

func _ReplyData_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ReplyData)
	switch tag {
	case 3: // msg.replyData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Reply)
		err := b.DecodeMessage(msg)
		m.Msg = &ReplyData_ReplyData{msg}
		return true, err
	case 4: // msg.replyContact
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ReplyContact)
		err := b.DecodeMessage(msg)
		m.Msg = &ReplyData_ReplyContact{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ReplyData_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ReplyData)
	// msg
	switch x := m.Msg.(type) {
	case *ReplyData_ReplyData:
		s := proto.Size(x.ReplyData)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ReplyData_ReplyContact:
		s := proto.Size(x.ReplyContact)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type WrapperMessage struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	SourceID string `protobuf:"bytes,2,opt,name=sourceID" json:"sourceID,omitempty"`
	// Types that are valid to be assigned to Msg:
	//	*WrapperMessage_M1
	//	*WrapperMessage_M2
	//	*WrapperMessage_M3
	//	*WrapperMessage_M4
	//	*WrapperMessage_M5
	//	*WrapperMessage_ReplyData
	Msg isWrapperMessage_Msg `protobuf_oneof:"msg"`
}

func (m *WrapperMessage) Reset()                    { *m = WrapperMessage{} }
func (m *WrapperMessage) String() string            { return proto.CompactTextString(m) }
func (*WrapperMessage) ProtoMessage()               {}
func (*WrapperMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isWrapperMessage_Msg interface {
	isWrapperMessage_Msg()
}

type WrapperMessage_M1 struct {
	M1 *RequestPing `protobuf:"bytes,3,opt,name=m1,oneof"`
}
type WrapperMessage_M2 struct {
	M2 *RequestContact `protobuf:"bytes,4,opt,name=m2,oneof"`
}
type WrapperMessage_M3 struct {
	M3 *RequestData `protobuf:"bytes,5,opt,name=m3,oneof"`
}
type WrapperMessage_M4 struct {
	M4 *Reply `protobuf:"bytes,6,opt,name=m4,oneof"`
}
type WrapperMessage_M5 struct {
	M5 *ReplyContact `protobuf:"bytes,7,opt,name=m5,oneof"`
}
type WrapperMessage_ReplyData struct {
	ReplyData *ReplyData `protobuf:"bytes,8,opt,name=replyData,oneof"`
}

func (*WrapperMessage_M1) isWrapperMessage_Msg()        {}
func (*WrapperMessage_M2) isWrapperMessage_Msg()        {}
func (*WrapperMessage_M3) isWrapperMessage_Msg()        {}
func (*WrapperMessage_M4) isWrapperMessage_Msg()        {}
func (*WrapperMessage_M5) isWrapperMessage_Msg()        {}
func (*WrapperMessage_ReplyData) isWrapperMessage_Msg() {}

func (m *WrapperMessage) GetMsg() isWrapperMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *WrapperMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *WrapperMessage) GetSourceID() string {
	if m != nil {
		return m.SourceID
	}
	return ""
}

func (m *WrapperMessage) GetM1() *RequestPing {
	if x, ok := m.GetMsg().(*WrapperMessage_M1); ok {
		return x.M1
	}
	return nil
}

func (m *WrapperMessage) GetM2() *RequestContact {
	if x, ok := m.GetMsg().(*WrapperMessage_M2); ok {
		return x.M2
	}
	return nil
}

func (m *WrapperMessage) GetM3() *RequestData {
	if x, ok := m.GetMsg().(*WrapperMessage_M3); ok {
		return x.M3
	}
	return nil
}

func (m *WrapperMessage) GetM4() *Reply {
	if x, ok := m.GetMsg().(*WrapperMessage_M4); ok {
		return x.M4
	}
	return nil
}

func (m *WrapperMessage) GetM5() *ReplyContact {
	if x, ok := m.GetMsg().(*WrapperMessage_M5); ok {
		return x.M5
	}
	return nil
}

func (m *WrapperMessage) GetReplyData() *ReplyData {
	if x, ok := m.GetMsg().(*WrapperMessage_ReplyData); ok {
		return x.ReplyData
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*WrapperMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _WrapperMessage_OneofMarshaler, _WrapperMessage_OneofUnmarshaler, _WrapperMessage_OneofSizer, []interface{}{
		(*WrapperMessage_M1)(nil),
		(*WrapperMessage_M2)(nil),
		(*WrapperMessage_M3)(nil),
		(*WrapperMessage_M4)(nil),
		(*WrapperMessage_M5)(nil),
		(*WrapperMessage_ReplyData)(nil),
	}
}

func _WrapperMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*WrapperMessage)
	// msg
	switch x := m.Msg.(type) {
	case *WrapperMessage_M1:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.M1); err != nil {
			return err
		}
	case *WrapperMessage_M2:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.M2); err != nil {
			return err
		}
	case *WrapperMessage_M3:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.M3); err != nil {
			return err
		}
	case *WrapperMessage_M4:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.M4); err != nil {
			return err
		}
	case *WrapperMessage_M5:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.M5); err != nil {
			return err
		}
	case *WrapperMessage_ReplyData:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReplyData); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("WrapperMessage.Msg has unexpected type %T", x)
	}
	return nil
}

func _WrapperMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*WrapperMessage)
	switch tag {
	case 3: // msg.m1
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestPing)
		err := b.DecodeMessage(msg)
		m.Msg = &WrapperMessage_M1{msg}
		return true, err
	case 4: // msg.m2
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestContact)
		err := b.DecodeMessage(msg)
		m.Msg = &WrapperMessage_M2{msg}
		return true, err
	case 5: // msg.m3
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestData)
		err := b.DecodeMessage(msg)
		m.Msg = &WrapperMessage_M3{msg}
		return true, err
	case 6: // msg.m4
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Reply)
		err := b.DecodeMessage(msg)
		m.Msg = &WrapperMessage_M4{msg}
		return true, err
	case 7: // msg.m5
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ReplyContact)
		err := b.DecodeMessage(msg)
		m.Msg = &WrapperMessage_M5{msg}
		return true, err
	case 8: // msg.replyData
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ReplyData)
		err := b.DecodeMessage(msg)
		m.Msg = &WrapperMessage_ReplyData{msg}
		return true, err
	default:
		return false, nil
	}
}

func _WrapperMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*WrapperMessage)
	// msg
	switch x := m.Msg.(type) {
	case *WrapperMessage_M1:
		s := proto.Size(x.M1)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WrapperMessage_M2:
		s := proto.Size(x.M2)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WrapperMessage_M3:
		s := proto.Size(x.M3)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WrapperMessage_M4:
		s := proto.Size(x.M4)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WrapperMessage_M5:
		s := proto.Size(x.M5)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WrapperMessage_ReplyData:
		s := proto.Size(x.ReplyData)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*RequestPing)(nil), "main.RequestPing")
	proto.RegisterType((*RequestContact)(nil), "main.RequestContact")
	proto.RegisterType((*RequestData)(nil), "main.RequestData")
	proto.RegisterType((*RequestStore)(nil), "main.RequestStore")
	proto.RegisterType((*Reply)(nil), "main.Reply")
	proto.RegisterType((*ReplyContact)(nil), "main.ReplyContact")
	proto.RegisterType((*ReplyContact_Contact)(nil), "main.ReplyContact.Contact")
	proto.RegisterType((*ReplyData)(nil), "main.ReplyData")
	proto.RegisterType((*WrapperMessage)(nil), "main.WrapperMessage")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x6d, 0xa7, 0xdf, 0xb7, 0x4b, 0xd5, 0x8b, 0xc8, 0x10, 0x58, 0x59, 0x46, 0x91, 0x85, 0x85,
	0x2c, 0x9b, 0xee, 0xca, 0xbe, 0xaa, 0x79, 0x68, 0x11, 0x51, 0xa2, 0xe0, 0xf3, 0xd8, 0x0c, 0x21,
	0xb0, 0x93, 0xc4, 0x99, 0xe9, 0x43, 0x7e, 0x91, 0x88, 0x7f, 0x52, 0x32, 0x4c, 0xd2, 0xa9, 0x29,
	0xee, 0x53, 0xe7, 0xce, 0x9c, 0x73, 0xee, 0x39, 0x87, 0x06, 0x20, 0xe5, 0x86, 0x87, 0x95, 0x2a,
	0x4d, 0x89, 0x63, 0xc9, 0xf3, 0x82, 0x9d, 0xc3, 0x32, 0x11, 0x3f, 0xf7, 0x42, 0x9b, 0x2f, 0x79,
	0x91, 0xe1, 0x0a, 0xc8, 0x36, 0xa5, 0xc3, 0x8b, 0xe1, 0xe5, 0x22, 0x21, 0xdb, 0x94, 0xdd, 0xc3,
	0xca, 0x3d, 0x7f, 0x28, 0x0b, 0xc3, 0x77, 0xe6, 0x5f, 0x04, 0xbe, 0x80, 0xa9, 0xe1, 0x2a, 0x13,
	0x86, 0x12, 0x7b, 0xe7, 0x26, 0x76, 0xdd, 0x09, 0xc7, 0xdc, 0xf0, 0x1e, 0xed, 0x29, 0x8c, 0x3e,
	0x8a, 0xda, 0x71, 0x9a, 0x23, 0x8b, 0xe1, 0xcc, 0x11, 0xbe, 0x9a, 0x52, 0x89, 0xc7, 0x19, 0x88,
	0x30, 0x6e, 0xb4, 0xe9, 0xc8, 0x5e, 0xd9, 0x33, 0xbb, 0x82, 0x49, 0x22, 0xaa, 0x87, 0xba, 0x47,
	0x47, 0x18, 0x37, 0xe1, 0x1d, 0xdf, 0x9e, 0xd9, 0xaf, 0x61, 0xb3, 0xb3, 0x7a, 0xa8, 0xbd, 0x70,
	0x79, 0x47, 0xca, 0x53, 0x7c, 0x0b, 0x73, 0xf7, 0xa4, 0x29, 0xb9, 0x18, 0x5d, 0x2e, 0xa3, 0x20,
	0x6c, 0x6a, 0x0b, 0x7d, 0x56, 0xe8, 0x7e, 0x93, 0x0e, 0x1b, 0x7c, 0x86, 0x99, 0xdf, 0x57, 0xdc,
	0xf9, 0x88, 0x91, 0xc2, 0xec, 0x5d, 0x9a, 0x2a, 0xa1, 0xb5, 0xb3, 0xd2, 0x8e, 0x18, 0xc0, 0x3c,
	0xce, 0xb5, 0xe1, 0xc5, 0x4e, 0xb8, 0x48, 0xdd, 0xcc, 0x7e, 0x0f, 0x61, 0x61, 0x77, 0xb6, 0x65,
	0x1e, 0xd9, 0x7c, 0x09, 0xa0, 0x84, 0xd9, 0xab, 0xe2, 0x5b, 0x5d, 0x09, 0x27, 0xeb, 0xdd, 0xe0,
	0x15, 0x2c, 0x54, 0x4b, 0xb6, 0xd2, 0xcb, 0x68, 0xe9, 0xe5, 0xd8, 0x0c, 0x92, 0xc3, 0x3b, 0xde,
	0xc3, 0x99, 0xf2, 0xd2, 0xd1, 0xb1, 0xc5, 0x63, 0x3f, 0xf7, 0x66, 0x90, 0x1c, 0x21, 0xdf, 0x4f,
	0x60, 0x24, 0x75, 0xc6, 0xfe, 0x10, 0x58, 0x7d, 0x57, 0xbc, 0xaa, 0x84, 0xfa, 0x24, 0xb4, 0xe6,
	0x99, 0xe8, 0x19, 0x0e, 0x60, 0xae, 0xcb, 0xbd, 0xda, 0x89, 0x6d, 0xec, 0xec, 0x76, 0x33, 0xbe,
	0x02, 0x22, 0x6f, 0x9c, 0xcb, 0x67, 0xed, 0xd6, 0xee, 0x1f, 0xba, 0x19, 0x24, 0x44, 0xde, 0xe0,
	0x1b, 0x20, 0x32, 0x72, 0xd6, 0x9e, 0x1f, 0x81, 0x0e, 0xe6, 0x88, 0x8c, 0xac, 0xd8, 0x9a, 0x4e,
	0x4e, 0x88, 0x35, 0x59, 0x2d, 0x68, 0x8d, 0xe7, 0x40, 0xe4, 0x2d, 0x9d, 0x9e, 0xea, 0x85, 0xc8,
	0x5b, 0x7c, 0x0d, 0x44, 0xde, 0xd1, 0xd9, 0x7f, 0x6a, 0x20, 0xf2, 0x0e, 0xaf, 0xfd, 0x8e, 0xe7,
	0x16, 0xfc, 0xc4, 0x03, 0xbb, 0x75, 0x07, 0x8c, 0x6b, 0xeb, 0xc7, 0xd4, 0x7e, 0x8d, 0xeb, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x98, 0x87, 0x76, 0x1e, 0x9b, 0x03, 0x00, 0x00,
}
