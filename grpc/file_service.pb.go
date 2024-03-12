// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: grpc/file_service.proto

package fileservice

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

type GetSingleFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
}

func (x *GetSingleFileRequest) Reset() {
	*x = GetSingleFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSingleFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSingleFileRequest) ProtoMessage() {}

func (x *GetSingleFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSingleFileRequest.ProtoReflect.Descriptor instead.
func (*GetSingleFileRequest) Descriptor() ([]byte, []int) {
	return file_grpc_file_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSingleFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type GetSingleFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *GetSingleFileResponse) Reset() {
	*x = GetSingleFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSingleFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSingleFileResponse) ProtoMessage() {}

func (x *GetSingleFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSingleFileResponse.ProtoReflect.Descriptor instead.
func (*GetSingleFileResponse) Descriptor() ([]byte, []int) {
	return file_grpc_file_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetSingleFileResponse) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type UploadSingleFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *UploadSingleFileRequest) Reset() {
	*x = UploadSingleFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadSingleFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadSingleFileRequest) ProtoMessage() {}

func (x *UploadSingleFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadSingleFileRequest.ProtoReflect.Descriptor instead.
func (*UploadSingleFileRequest) Descriptor() ([]byte, []int) {
	return file_grpc_file_service_proto_rawDescGZIP(), []int{2}
}

func (x *UploadSingleFileRequest) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type UploadSingleFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
}

func (x *UploadSingleFileResponse) Reset() {
	*x = UploadSingleFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadSingleFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadSingleFileResponse) ProtoMessage() {}

func (x *UploadSingleFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadSingleFileResponse.ProtoReflect.Descriptor instead.
func (*UploadSingleFileResponse) Descriptor() ([]byte, []int) {
	return file_grpc_file_service_proto_rawDescGZIP(), []int{3}
}

func (x *UploadSingleFileResponse) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DeleteSingleFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId string `protobuf:"bytes,1,opt,name=fileId,proto3" json:"fileId,omitempty"`
}

func (x *DeleteSingleFileRequest) Reset() {
	*x = DeleteSingleFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSingleFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSingleFileRequest) ProtoMessage() {}

func (x *DeleteSingleFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSingleFileRequest.ProtoReflect.Descriptor instead.
func (*DeleteSingleFileRequest) Descriptor() ([]byte, []int) {
	return file_grpc_file_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSingleFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DeleteSingleFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Done bool `protobuf:"varint,1,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *DeleteSingleFileResponse) Reset() {
	*x = DeleteSingleFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSingleFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSingleFileResponse) ProtoMessage() {}

func (x *DeleteSingleFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSingleFileResponse.ProtoReflect.Descriptor instead.
func (*DeleteSingleFileResponse) Descriptor() ([]byte, []int) {
	return file_grpc_file_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSingleFileResponse) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

var File_grpc_file_service_proto protoreflect.FileDescriptor

var file_grpc_file_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x2e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x17, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x32, 0x0a, 0x18, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x18, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x32, 0xad, 0x02, 0x0a, 0x0b, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x64, 0x0a, 0x2c, 0x63, 0x6f, 0x6d,
	0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5a, 0x34, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x66, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_file_service_proto_rawDescOnce sync.Once
	file_grpc_file_service_proto_rawDescData = file_grpc_file_service_proto_rawDesc
)

func file_grpc_file_service_proto_rawDescGZIP() []byte {
	file_grpc_file_service_proto_rawDescOnce.Do(func() {
		file_grpc_file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_file_service_proto_rawDescData)
	})
	return file_grpc_file_service_proto_rawDescData
}

var file_grpc_file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_file_service_proto_goTypes = []interface{}{
	(*GetSingleFileRequest)(nil),     // 0: fileservice.GetSingleFileRequest
	(*GetSingleFileResponse)(nil),    // 1: fileservice.GetSingleFileResponse
	(*UploadSingleFileRequest)(nil),  // 2: fileservice.UploadSingleFileRequest
	(*UploadSingleFileResponse)(nil), // 3: fileservice.UploadSingleFileResponse
	(*DeleteSingleFileRequest)(nil),  // 4: fileservice.DeleteSingleFileRequest
	(*DeleteSingleFileResponse)(nil), // 5: fileservice.DeleteSingleFileResponse
}
var file_grpc_file_service_proto_depIdxs = []int32{
	0, // 0: fileservice.FileService.GetSingleFile:input_type -> fileservice.GetSingleFileRequest
	2, // 1: fileservice.FileService.UploadSingleFile:input_type -> fileservice.UploadSingleFileRequest
	4, // 2: fileservice.FileService.DeleteSingleFile:input_type -> fileservice.DeleteSingleFileRequest
	1, // 3: fileservice.FileService.GetSingleFile:output_type -> fileservice.GetSingleFileResponse
	3, // 4: fileservice.FileService.UploadSingleFile:output_type -> fileservice.UploadSingleFileResponse
	5, // 5: fileservice.FileService.DeleteSingleFile:output_type -> fileservice.DeleteSingleFileResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_file_service_proto_init() }
func file_grpc_file_service_proto_init() {
	if File_grpc_file_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSingleFileRequest); i {
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
		file_grpc_file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSingleFileResponse); i {
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
		file_grpc_file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadSingleFileRequest); i {
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
		file_grpc_file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadSingleFileResponse); i {
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
		file_grpc_file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSingleFileRequest); i {
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
		file_grpc_file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSingleFileResponse); i {
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
			RawDescriptor: file_grpc_file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_file_service_proto_goTypes,
		DependencyIndexes: file_grpc_file_service_proto_depIdxs,
		MessageInfos:      file_grpc_file_service_proto_msgTypes,
	}.Build()
	File_grpc_file_service_proto = out.File
	file_grpc_file_service_proto_rawDesc = nil
	file_grpc_file_service_proto_goTypes = nil
	file_grpc_file_service_proto_depIdxs = nil
}
