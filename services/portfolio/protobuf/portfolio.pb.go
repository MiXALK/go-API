// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: portfolio.proto

package portfolio

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreatePortfolioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
}

func (x *CreatePortfolioRequest) Reset() {
	*x = CreatePortfolioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portfolio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePortfolioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePortfolioRequest) ProtoMessage() {}

func (x *CreatePortfolioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_portfolio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePortfolioRequest.ProtoReflect.Descriptor instead.
func (*CreatePortfolioRequest) Descriptor() ([]byte, []int) {
	return file_portfolio_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePortfolioRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreatePortfolioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=Status,json=status,proto3" json:"Status,omitempty"`
	ID     string `protobuf:"bytes,2,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
}

func (x *CreatePortfolioResponse) Reset() {
	*x = CreatePortfolioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portfolio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePortfolioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePortfolioResponse) ProtoMessage() {}

func (x *CreatePortfolioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_portfolio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePortfolioResponse.ProtoReflect.Descriptor instead.
func (*CreatePortfolioResponse) Descriptor() ([]byte, []int) {
	return file_portfolio_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePortfolioResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreatePortfolioResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type Portfolio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
}

func (x *Portfolio) Reset() {
	*x = Portfolio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portfolio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Portfolio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Portfolio) ProtoMessage() {}

func (x *Portfolio) ProtoReflect() protoreflect.Message {
	mi := &file_portfolio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Portfolio.ProtoReflect.Descriptor instead.
func (*Portfolio) Descriptor() ([]byte, []int) {
	return file_portfolio_proto_rawDescGZIP(), []int{2}
}

func (x *Portfolio) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Portfolio) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_portfolio_proto protoreflect.FileDescriptor

var file_portfolio_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x22, 0x2c, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x44, 0x22, 0x2f, 0x0a,
	0x09, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x65,
	0x0a, 0x10, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x70,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x66, 0x6f, 0x6c, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x70, 0x6f, 0x72, 0x74, 0x66,
	0x6f, 0x6c, 0x69, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_portfolio_proto_rawDescOnce sync.Once
	file_portfolio_proto_rawDescData = file_portfolio_proto_rawDesc
)

func file_portfolio_proto_rawDescGZIP() []byte {
	file_portfolio_proto_rawDescOnce.Do(func() {
		file_portfolio_proto_rawDescData = protoimpl.X.CompressGZIP(file_portfolio_proto_rawDescData)
	})
	return file_portfolio_proto_rawDescData
}

var file_portfolio_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_portfolio_proto_goTypes = []interface{}{
	(*CreatePortfolioRequest)(nil),  // 0: portfolio.CreatePortfolioRequest
	(*CreatePortfolioResponse)(nil), // 1: portfolio.CreatePortfolioResponse
	(*Portfolio)(nil),               // 2: portfolio.Portfolio
}
var file_portfolio_proto_depIdxs = []int32{
	0, // 0: portfolio.PortfolioService.Create:input_type -> portfolio.CreatePortfolioRequest
	1, // 1: portfolio.PortfolioService.Create:output_type -> portfolio.CreatePortfolioResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_portfolio_proto_init() }
func file_portfolio_proto_init() {
	if File_portfolio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_portfolio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePortfolioRequest); i {
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
		file_portfolio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePortfolioResponse); i {
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
		file_portfolio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Portfolio); i {
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
			RawDescriptor: file_portfolio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_portfolio_proto_goTypes,
		DependencyIndexes: file_portfolio_proto_depIdxs,
		MessageInfos:      file_portfolio_proto_msgTypes,
	}.Build()
	File_portfolio_proto = out.File
	file_portfolio_proto_rawDesc = nil
	file_portfolio_proto_goTypes = nil
	file_portfolio_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PortfolioServiceClient is the client API for PortfolioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortfolioServiceClient interface {
	Create(ctx context.Context, in *CreatePortfolioRequest, opts ...grpc.CallOption) (*CreatePortfolioResponse, error)
}

type portfolioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortfolioServiceClient(cc grpc.ClientConnInterface) PortfolioServiceClient {
	return &portfolioServiceClient{cc}
}

func (c *portfolioServiceClient) Create(ctx context.Context, in *CreatePortfolioRequest, opts ...grpc.CallOption) (*CreatePortfolioResponse, error) {
	out := new(CreatePortfolioResponse)
	err := c.cc.Invoke(ctx, "/portfolio.PortfolioService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortfolioServiceServer is the server API for PortfolioService service.
type PortfolioServiceServer interface {
	Create(context.Context, *CreatePortfolioRequest) (*CreatePortfolioResponse, error)
}

// UnimplementedPortfolioServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPortfolioServiceServer struct {
}

func (*UnimplementedPortfolioServiceServer) Create(context.Context, *CreatePortfolioRequest) (*CreatePortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterPortfolioServiceServer(s *grpc.Server, srv PortfolioServiceServer) {
	s.RegisterService(&_PortfolioService_serviceDesc, srv)
}

func _PortfolioService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePortfolioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortfolioServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/portfolio.PortfolioService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortfolioServiceServer).Create(ctx, req.(*CreatePortfolioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortfolioService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "portfolio.PortfolioService",
	HandlerType: (*PortfolioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PortfolioService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portfolio.proto",
}
