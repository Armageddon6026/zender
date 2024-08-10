// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.7
// source: xsaber.proto

package gcontroller

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

type LoginServiceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid             string             `protobuf:"bytes,1,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	Name             string             `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Ip               string             `protobuf:"bytes,3,opt,name=Ip,proto3" json:"Ip,omitempty"`
	ScanFiles        []*ScanFile        `protobuf:"bytes,4,rep,name=ScanFiles,proto3" json:"ScanFiles,omitempty"`
	ScanApplications []*ScanApplication `protobuf:"bytes,5,rep,name=ScanApplications,proto3" json:"ScanApplications,omitempty"`
}

func (x *LoginServiceInfo) Reset() {
	*x = LoginServiceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xsaber_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginServiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginServiceInfo) ProtoMessage() {}

func (x *LoginServiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_xsaber_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginServiceInfo.ProtoReflect.Descriptor instead.
func (*LoginServiceInfo) Descriptor() ([]byte, []int) {
	return file_xsaber_proto_rawDescGZIP(), []int{0}
}

func (x *LoginServiceInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *LoginServiceInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginServiceInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *LoginServiceInfo) GetScanFiles() []*ScanFile {
	if x != nil {
		return x.ScanFiles
	}
	return nil
}

func (x *LoginServiceInfo) GetScanApplications() []*ScanApplication {
	if x != nil {
		return x.ScanApplications
	}
	return nil
}

type ScanFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	DataCount int32  `protobuf:"varint,2,opt,name=DataCount,proto3" json:"DataCount,omitempty"`
	LastTime  int64  `protobuf:"varint,3,opt,name=LastTime,proto3" json:"LastTime,omitempty"`
}

func (x *ScanFile) Reset() {
	*x = ScanFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xsaber_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanFile) ProtoMessage() {}

func (x *ScanFile) ProtoReflect() protoreflect.Message {
	mi := &file_xsaber_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanFile.ProtoReflect.Descriptor instead.
func (*ScanFile) Descriptor() ([]byte, []int) {
	return file_xsaber_proto_rawDescGZIP(), []int{1}
}

func (x *ScanFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ScanFile) GetDataCount() int32 {
	if x != nil {
		return x.DataCount
	}
	return 0
}

func (x *ScanFile) GetLastTime() int64 {
	if x != nil {
		return x.LastTime
	}
	return 0
}

type ScanApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	LastTime int64  `protobuf:"varint,2,opt,name=LastTime,proto3" json:"LastTime,omitempty"`
}

func (x *ScanApplication) Reset() {
	*x = ScanApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xsaber_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanApplication) ProtoMessage() {}

func (x *ScanApplication) ProtoReflect() protoreflect.Message {
	mi := &file_xsaber_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanApplication.ProtoReflect.Descriptor instead.
func (*ScanApplication) Descriptor() ([]byte, []int) {
	return file_xsaber_proto_rawDescGZIP(), []int{2}
}

func (x *ScanApplication) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScanApplication) GetLastTime() int64 {
	if x != nil {
		return x.LastTime
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xsaber_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_xsaber_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_xsaber_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_xsaber_proto protoreflect.FileDescriptor

var file_xsaber_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x78, 0x73, 0x61, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x67, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0xc9, 0x01, 0x0a, 0x10,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x55, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x70, 0x12, 0x33, 0x0a, 0x09, 0x53, 0x63, 0x61, 0x6e,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x09, 0x53, 0x63, 0x61, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x48, 0x0a,
	0x10, 0x53, 0x63, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x53, 0x63, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x58, 0x0a, 0x08, 0x53, 0x63, 0x61, 0x6e, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x41, 0x0a, 0x0f, 0x53, 0x63, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb0,
	0x01, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x75, 0x6e, 0x63, 0x12,
	0x4e, 0x0a, 0x16, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x67, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x67, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x50, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x2e, 0x67, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x15, 0x2e, 0x67, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x42, 0x23, 0x5a, 0x21, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xsaber_proto_rawDescOnce sync.Once
	file_xsaber_proto_rawDescData = file_xsaber_proto_rawDesc
)

func file_xsaber_proto_rawDescGZIP() []byte {
	file_xsaber_proto_rawDescOnce.Do(func() {
		file_xsaber_proto_rawDescData = protoimpl.X.CompressGZIP(file_xsaber_proto_rawDescData)
	})
	return file_xsaber_proto_rawDescData
}

var file_xsaber_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_xsaber_proto_goTypes = []interface{}{
	(*LoginServiceInfo)(nil), // 0: gcontroller.LoginServiceInfo
	(*ScanFile)(nil),         // 1: gcontroller.ScanFile
	(*ScanApplication)(nil),  // 2: gcontroller.ScanApplication
	(*Response)(nil),         // 3: gcontroller.Response
}
var file_xsaber_proto_depIdxs = []int32{
	1, // 0: gcontroller.LoginServiceInfo.ScanFiles:type_name -> gcontroller.ScanFile
	2, // 1: gcontroller.LoginServiceInfo.ScanApplications:type_name -> gcontroller.ScanApplication
	0, // 2: gcontroller.ServicesFunc.InsertLoginServiceInfo:input_type -> gcontroller.LoginServiceInfo
	0, // 3: gcontroller.ServicesFunc.UpdateLoginServiceInfo:input_type -> gcontroller.LoginServiceInfo
	3, // 4: gcontroller.ServicesFunc.InsertLoginServiceInfo:output_type -> gcontroller.Response
	3, // 5: gcontroller.ServicesFunc.UpdateLoginServiceInfo:output_type -> gcontroller.Response
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_xsaber_proto_init() }
func file_xsaber_proto_init() {
	if File_xsaber_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xsaber_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginServiceInfo); i {
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
		file_xsaber_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanFile); i {
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
		file_xsaber_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanApplication); i {
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
		file_xsaber_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_xsaber_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xsaber_proto_goTypes,
		DependencyIndexes: file_xsaber_proto_depIdxs,
		MessageInfos:      file_xsaber_proto_msgTypes,
	}.Build()
	File_xsaber_proto = out.File
	file_xsaber_proto_rawDesc = nil
	file_xsaber_proto_goTypes = nil
	file_xsaber_proto_depIdxs = nil
}
