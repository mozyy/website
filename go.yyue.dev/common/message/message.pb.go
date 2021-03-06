// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.10.1
// source: common/message.proto

package message

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Status int32

const (
	Status_Success Status = 0
	Status_Info    Status = 1
	Status_Warning Status = 2
	Status_Error   Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Success",
		1: "Info",
		2: "Warning",
		3: "Error",
	}
	Status_value = map[string]int32{
		"Success": 0,
		"Info":    1,
		"Warning": 2,
		"Error":   3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_common_message_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_common_message_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_common_message_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State   Status `protobuf:"varint,1,opt,name=state,proto3,enum=message.Status" json:"state,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_common_message_proto_msgTypes[0]
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
	return file_common_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetState() Status {
	if x != nil {
		return x.State
	}
	return Status_Success
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_common_message_proto protoreflect.FileDescriptor

var file_common_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x4a, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x37, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x03, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x6f, 0x2e, 0x79, 0x79, 0x75, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_message_proto_rawDescOnce sync.Once
	file_common_message_proto_rawDescData = file_common_message_proto_rawDesc
)

func file_common_message_proto_rawDescGZIP() []byte {
	file_common_message_proto_rawDescOnce.Do(func() {
		file_common_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_message_proto_rawDescData)
	})
	return file_common_message_proto_rawDescData
}

var file_common_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_message_proto_goTypes = []interface{}{
	(Status)(0),     // 0: message.Status
	(*Message)(nil), // 1: message.Message
}
var file_common_message_proto_depIdxs = []int32{
	0, // 0: message.Message.state:type_name -> message.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_message_proto_init() }
func file_common_message_proto_init() {
	if File_common_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_message_proto_goTypes,
		DependencyIndexes: file_common_message_proto_depIdxs,
		EnumInfos:         file_common_message_proto_enumTypes,
		MessageInfos:      file_common_message_proto_msgTypes,
	}.Build()
	File_common_message_proto = out.File
	file_common_message_proto_rawDesc = nil
	file_common_message_proto_goTypes = nil
	file_common_message_proto_depIdxs = nil
}
