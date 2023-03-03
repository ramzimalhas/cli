// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: jupyter/jupyter_server_host_service.v1.proto

package jupyter

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

type GetRunningServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRunningServerRequest) Reset() {
	*x = GetRunningServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jupyter_jupyter_server_host_service_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRunningServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRunningServerRequest) ProtoMessage() {}

func (x *GetRunningServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jupyter_jupyter_server_host_service_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRunningServerRequest.ProtoReflect.Descriptor instead.
func (*GetRunningServerRequest) Descriptor() ([]byte, []int) {
	return file_jupyter_jupyter_server_host_service_v1_proto_rawDescGZIP(), []int{0}
}

type GetRunningServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    bool   `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Port      string `protobuf:"bytes,3,opt,name=Port,proto3" json:"Port,omitempty"`
	ServerUrl string `protobuf:"bytes,4,opt,name=ServerUrl,proto3" json:"ServerUrl,omitempty"`
}

func (x *GetRunningServerResponse) Reset() {
	*x = GetRunningServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jupyter_jupyter_server_host_service_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRunningServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRunningServerResponse) ProtoMessage() {}

func (x *GetRunningServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jupyter_jupyter_server_host_service_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRunningServerResponse.ProtoReflect.Descriptor instead.
func (*GetRunningServerResponse) Descriptor() ([]byte, []int) {
	return file_jupyter_jupyter_server_host_service_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GetRunningServerResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *GetRunningServerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetRunningServerResponse) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *GetRunningServerResponse) GetServerUrl() string {
	if x != nil {
		return x.ServerUrl
	}
	return ""
}

var File_jupyter_jupyter_server_host_service_v1_proto protoreflect.FileDescriptor

var file_jupyter_jupyter_server_host_service_v1_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6a, 0x75, 0x70, 0x79, 0x74, 0x65, 0x72, 0x2f, 0x6a, 0x75, 0x70, 0x79, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b,
	0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e,
	0x4a, 0x75, 0x70, 0x79, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x6f, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x19, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x32, 0xb5, 0x01, 0x0a, 0x11, 0x4a, 0x75, 0x70, 0x79, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x9f, 0x01, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x44, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x47,
	0x72, 0x70, 0x63, 0x2e, 0x4a, 0x75, 0x70, 0x79, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x75, 0x70, 0x79, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x6a, 0x75, 0x70, 0x79, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_jupyter_jupyter_server_host_service_v1_proto_rawDescOnce sync.Once
	file_jupyter_jupyter_server_host_service_v1_proto_rawDescData = file_jupyter_jupyter_server_host_service_v1_proto_rawDesc
)

func file_jupyter_jupyter_server_host_service_v1_proto_rawDescGZIP() []byte {
	file_jupyter_jupyter_server_host_service_v1_proto_rawDescOnce.Do(func() {
		file_jupyter_jupyter_server_host_service_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_jupyter_jupyter_server_host_service_v1_proto_rawDescData)
	})
	return file_jupyter_jupyter_server_host_service_v1_proto_rawDescData
}

var file_jupyter_jupyter_server_host_service_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_jupyter_jupyter_server_host_service_v1_proto_goTypes = []interface{}{
	(*GetRunningServerRequest)(nil),  // 0: Codespaces.Grpc.JupyterServerHostService.v1.GetRunningServerRequest
	(*GetRunningServerResponse)(nil), // 1: Codespaces.Grpc.JupyterServerHostService.v1.GetRunningServerResponse
}
var file_jupyter_jupyter_server_host_service_v1_proto_depIdxs = []int32{
	0, // 0: Codespaces.Grpc.JupyterServerHostService.v1.JupyterServerHost.GetRunningServer:input_type -> Codespaces.Grpc.JupyterServerHostService.v1.GetRunningServerRequest
	1, // 1: Codespaces.Grpc.JupyterServerHostService.v1.JupyterServerHost.GetRunningServer:output_type -> Codespaces.Grpc.JupyterServerHostService.v1.GetRunningServerResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_jupyter_jupyter_server_host_service_v1_proto_init() }
func file_jupyter_jupyter_server_host_service_v1_proto_init() {
	if File_jupyter_jupyter_server_host_service_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jupyter_jupyter_server_host_service_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRunningServerRequest); i {
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
		file_jupyter_jupyter_server_host_service_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRunningServerResponse); i {
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
			RawDescriptor: file_jupyter_jupyter_server_host_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jupyter_jupyter_server_host_service_v1_proto_goTypes,
		DependencyIndexes: file_jupyter_jupyter_server_host_service_v1_proto_depIdxs,
		MessageInfos:      file_jupyter_jupyter_server_host_service_v1_proto_msgTypes,
	}.Build()
	File_jupyter_jupyter_server_host_service_v1_proto = out.File
	file_jupyter_jupyter_server_host_service_v1_proto_rawDesc = nil
	file_jupyter_jupyter_server_host_service_v1_proto_goTypes = nil
	file_jupyter_jupyter_server_host_service_v1_proto_depIdxs = nil
}
