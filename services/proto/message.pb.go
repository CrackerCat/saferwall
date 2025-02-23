// Copyright 2018 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: message.proto

package saferwall

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PayloadKind represents the kind of message.
type Message_PayloadKind int32

const (
	Message_DBUPDATE Message_PayloadKind = 0
	Message_DBCREATE Message_PayloadKind = 1
	Message_UPLOAD   Message_PayloadKind = 2
)

// Enum value maps for Message_PayloadKind.
var (
	Message_PayloadKind_name = map[int32]string{
		0: "DBUPDATE",
		1: "DBCREATE",
		2: "UPLOAD",
	}
	Message_PayloadKind_value = map[string]int32{
		"DBUPDATE": 0,
		"DBCREATE": 1,
		"UPLOAD":   2,
	}
)

func (x Message_PayloadKind) Enum() *Message_PayloadKind {
	p := new(Message_PayloadKind)
	*p = x
	return p
}

func (x Message_PayloadKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_PayloadKind) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (Message_PayloadKind) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x Message_PayloadKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Message_PayloadKind.Descriptor instead.
func (Message_PayloadKind) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0, 0}
}

// The services message definition.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SHA256 of the binary file.
	Sha256  string             `protobuf:"bytes,1,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Payload []*Message_Payload `protobuf:"bytes,2,rep,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (x *Message) GetPayload() []*Message_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Payload represents the body of the message.
type Message_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kind represents the type of payload.
	Kind Message_PayloadKind `protobuf:"varint,1,opt,name=kind,proto3,enum=service.Message_PayloadKind" json:"kind,omitempty"`
	// The key to use to write the payload, can be either an
	// object stortage key or a DB document key.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Path represents the DB path where to write the payload
	// when the message kind is a DBUPDATE.
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// The raw body.
	Body []byte `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message_Payload) Reset() {
	*x = Message_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_Payload) ProtoMessage() {}

func (x *Message_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_Payload.ProtoReflect.Descriptor instead.
func (*Message_Payload) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Message_Payload) GetKind() Message_PayloadKind {
	if x != nil {
		return x.Kind
	}
	return Message_DBUPDATE
}

func (x *Message_Payload) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Message_Payload) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Message_Payload) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x83, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12, 0x32, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x1a, 0x75, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x35, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x42, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x42, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x02, 0x42, 0x20,
	0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x66,
	0x65, 0x72, 0x77, 0x61, 0x6c, 0x6c, 0x2f, 0x73, 0x61, 0x66, 0x65, 0x72, 0x77, 0x61, 0x6c, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_proto_goTypes = []interface{}{
	(Message_PayloadKind)(0), // 0: service.Message.PayloadKind
	(*Message)(nil),          // 1: service.Message
	(*Message_Payload)(nil),  // 2: service.Message.Payload
}
var file_message_proto_depIdxs = []int32{
	2, // 0: service.Message.payload:type_name -> service.Message.Payload
	0, // 1: service.Message.Payload.kind:type_name -> service.Message.PayloadKind
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_Payload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
