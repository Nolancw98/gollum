// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	SerializedMessageData
	SerializedMessage
*/
package core

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

type SerializedMessageData struct {
	Data []byte `protobuf:"bytes,1,req,name=Data" json:"Data,omitempty"`
	// Field 2 is used for v0.5.x metadata
	Metadata         []byte `protobuf:"bytes,3,opt,name=Metadata" json:"Metadata,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SerializedMessageData) Reset()                    { *m = SerializedMessageData{} }
func (m *SerializedMessageData) String() string            { return proto.CompactTextString(m) }
func (*SerializedMessageData) ProtoMessage()               {}
func (*SerializedMessageData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SerializedMessageData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SerializedMessageData) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type SerializedMessage struct {
	StreamID         *uint64                `protobuf:"varint,1,req,name=StreamID" json:"StreamID,omitempty"`
	Data             *SerializedMessageData `protobuf:"bytes,2,req,name=Data" json:"Data,omitempty"`
	PrevStreamID     *uint64                `protobuf:"varint,3,opt,name=PrevStreamID" json:"PrevStreamID,omitempty"`
	OrigStreamID     *uint64                `protobuf:"varint,4,opt,name=OrigStreamID" json:"OrigStreamID,omitempty"`
	Timestamp        *int64                 `protobuf:"varint,5,opt,name=Timestamp" json:"Timestamp,omitempty"`
	Original         *SerializedMessageData `protobuf:"bytes,6,opt,name=Original" json:"Original,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *SerializedMessage) Reset()                    { *m = SerializedMessage{} }
func (m *SerializedMessage) String() string            { return proto.CompactTextString(m) }
func (*SerializedMessage) ProtoMessage()               {}
func (*SerializedMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SerializedMessage) GetStreamID() uint64 {
	if m != nil && m.StreamID != nil {
		return *m.StreamID
	}
	return 0
}

func (m *SerializedMessage) GetData() *SerializedMessageData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SerializedMessage) GetPrevStreamID() uint64 {
	if m != nil && m.PrevStreamID != nil {
		return *m.PrevStreamID
	}
	return 0
}

func (m *SerializedMessage) GetOrigStreamID() uint64 {
	if m != nil && m.OrigStreamID != nil {
		return *m.OrigStreamID
	}
	return 0
}

func (m *SerializedMessage) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *SerializedMessage) GetOriginal() *SerializedMessageData {
	if m != nil {
		return m.Original
	}
	return nil
}

func init() {
	proto.RegisterType((*SerializedMessageData)(nil), "serializedMessageData")
	proto.RegisterType((*SerializedMessage)(nil), "serializedMessage")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x72, 0xe7, 0x12, 0x2d, 0x4e, 0x2d,
	0xca, 0x4c, 0xcc, 0xc9, 0xac, 0x4a, 0x4d, 0xf1, 0x85, 0x48, 0xb9, 0x24, 0x96, 0x24, 0x0a, 0x09,
	0x71, 0xb1, 0x80, 0x68, 0x09, 0x46, 0x05, 0x26, 0x0d, 0x9e, 0x20, 0x30, 0x5b, 0x48, 0x8a, 0x8b,
	0xc3, 0x37, 0xb5, 0x24, 0x31, 0x05, 0x24, 0xce, 0xac, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0xe7, 0x2b,
	0x7d, 0x65, 0xe4, 0x12, 0xc4, 0x30, 0x09, 0xa4, 0x23, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0xd7, 0xd3,
	0x05, 0x6c, 0x12, 0x4b, 0x10, 0x9c, 0x2f, 0xa4, 0x05, 0xb5, 0x81, 0x49, 0x81, 0x49, 0x83, 0xdb,
	0x48, 0x4c, 0x0f, 0xab, 0x3b, 0xa0, 0x36, 0x2b, 0x71, 0xf1, 0x04, 0x14, 0xa5, 0x96, 0xc1, 0xcd,
	0x02, 0xd9, 0xce, 0x12, 0x84, 0x22, 0x06, 0x52, 0xe3, 0x5f, 0x94, 0x99, 0x0e, 0x57, 0xc3, 0x02,
	0x51, 0x83, 0x2c, 0x26, 0x24, 0xc3, 0xc5, 0x19, 0x92, 0x99, 0x9b, 0x5a, 0x5c, 0x92, 0x98, 0x5b,
	0x20, 0xc1, 0xaa, 0xc0, 0xa8, 0xc1, 0x1c, 0x84, 0x10, 0x10, 0x32, 0xe2, 0xe2, 0x00, 0xa9, 0xce,
	0xcc, 0x4b, 0xcc, 0x91, 0x60, 0x53, 0x60, 0xc4, 0xe3, 0x2a, 0xb8, 0x3a, 0x27, 0xb6, 0x28, 0x96,
	0xe4, 0xfc, 0xa2, 0x54, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xe9, 0xdd, 0x67, 0x58, 0x01,
	0x00, 0x00,
}
