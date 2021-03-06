// Code generated by protoc-gen-go. DO NOT EDIT.
// source: my_package.proto

package my_package

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Document_Type int32

const (
	Document_IMAGE Document_Type = 0
	Document_TEXT  Document_Type = 1
	Document_VIDEO Document_Type = 2
)

var Document_Type_name = map[int32]string{
	0: "IMAGE",
	1: "TEXT",
	2: "VIDEO",
}

var Document_Type_value = map[string]int32{
	"IMAGE": 0,
	"TEXT":  1,
	"VIDEO": 2,
}

func (x Document_Type) String() string {
	return proto.EnumName(Document_Type_name, int32(x))
}

func (Document_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f2e0d9d8de006ff, []int{4, 0}
}

type StoreRequest struct {
	Document             *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *StoreRequest) Reset()         { *m = StoreRequest{} }
func (m *StoreRequest) String() string { return proto.CompactTextString(m) }
func (*StoreRequest) ProtoMessage()    {}
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f2e0d9d8de006ff, []int{0}
}

func (m *StoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreRequest.Unmarshal(m, b)
}
func (m *StoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreRequest.Marshal(b, m, deterministic)
}
func (m *StoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreRequest.Merge(m, src)
}
func (m *StoreRequest) XXX_Size() int {
	return xxx_messageInfo_StoreRequest.Size(m)
}
func (m *StoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreRequest proto.InternalMessageInfo

func (m *StoreRequest) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type StoreResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreResponse) Reset()         { *m = StoreResponse{} }
func (m *StoreResponse) String() string { return proto.CompactTextString(m) }
func (*StoreResponse) ProtoMessage()    {}
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f2e0d9d8de006ff, []int{1}
}

func (m *StoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResponse.Unmarshal(m, b)
}
func (m *StoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResponse.Marshal(b, m, deterministic)
}
func (m *StoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResponse.Merge(m, src)
}
func (m *StoreResponse) XXX_Size() int {
	return xxx_messageInfo_StoreResponse.Size(m)
}
func (m *StoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResponse proto.InternalMessageInfo

type FindRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRequest) Reset()         { *m = FindRequest{} }
func (m *FindRequest) String() string { return proto.CompactTextString(m) }
func (*FindRequest) ProtoMessage()    {}
func (*FindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f2e0d9d8de006ff, []int{2}
}

func (m *FindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRequest.Unmarshal(m, b)
}
func (m *FindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRequest.Marshal(b, m, deterministic)
}
func (m *FindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRequest.Merge(m, src)
}
func (m *FindRequest) XXX_Size() int {
	return xxx_messageInfo_FindRequest.Size(m)
}
func (m *FindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRequest proto.InternalMessageInfo

type FindResponse struct {
	Document             *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
	Error                string    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FindResponse) Reset()         { *m = FindResponse{} }
func (m *FindResponse) String() string { return proto.CompactTextString(m) }
func (*FindResponse) ProtoMessage()    {}
func (*FindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f2e0d9d8de006ff, []int{3}
}

func (m *FindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResponse.Unmarshal(m, b)
}
func (m *FindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResponse.Marshal(b, m, deterministic)
}
func (m *FindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResponse.Merge(m, src)
}
func (m *FindResponse) XXX_Size() int {
	return xxx_messageInfo_FindResponse.Size(m)
}
func (m *FindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindResponse proto.InternalMessageInfo

func (m *FindResponse) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

func (m *FindResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Document struct {
	Title                string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type                 Document_Type `protobuf:"varint,2,opt,name=type,proto3,enum=my_package.Document_Type" json:"type,omitempty"`
	Description          string        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f2e0d9d8de006ff, []int{4}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Document) GetType() Document_Type {
	if m != nil {
		return m.Type
	}
	return Document_IMAGE
}

func (m *Document) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("my_package.Document_Type", Document_Type_name, Document_Type_value)
	proto.RegisterType((*StoreRequest)(nil), "my_package.StoreRequest")
	proto.RegisterType((*StoreResponse)(nil), "my_package.StoreResponse")
	proto.RegisterType((*FindRequest)(nil), "my_package.FindRequest")
	proto.RegisterType((*FindResponse)(nil), "my_package.FindResponse")
	proto.RegisterType((*Document)(nil), "my_package.Document")
}

func init() { proto.RegisterFile("my_package.proto", fileDescriptor_8f2e0d9d8de006ff) }

var fileDescriptor_8f2e0d9d8de006ff = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x4f, 0x4b, 0xc3, 0x30,
	0x1c, 0x5d, 0x66, 0x27, 0xeb, 0xaf, 0x9b, 0x96, 0x30, 0xb0, 0xdb, 0xa9, 0xe4, 0x20, 0xbb, 0x58,
	0x46, 0x3d, 0x2b, 0x0a, 0xad, 0xb2, 0xc3, 0x50, 0x6a, 0x19, 0xde, 0x64, 0xb6, 0x3f, 0xa4, 0xe8,
	0x9a, 0x98, 0x66, 0x87, 0x7e, 0x05, 0xbf, 0x81, 0xdf, 0x56, 0x9a, 0x56, 0x0d, 0xb8, 0x8b, 0xc7,
	0xf7, 0xe7, 0xf7, 0x5e, 0x78, 0x01, 0x77, 0x5b, 0x3f, 0x89, 0x4d, 0xf6, 0xba, 0x79, 0xc1, 0x40,
	0x48, 0xae, 0x38, 0x85, 0x5f, 0x86, 0x5d, 0xc1, 0xe8, 0x41, 0x71, 0x89, 0x09, 0xbe, 0xef, 0xb0,
	0x52, 0x74, 0x01, 0xc3, 0x9c, 0x67, 0xbb, 0x2d, 0x96, 0xca, 0x23, 0x3e, 0x99, 0x3b, 0xe1, 0x24,
	0x30, 0x02, 0xa2, 0x4e, 0x4b, 0x7e, 0x5c, 0xec, 0x18, 0xc6, 0x5d, 0x42, 0x25, 0x78, 0x59, 0x21,
	0x1b, 0x83, 0x73, 0x53, 0x94, 0x79, 0x97, 0xc8, 0xd6, 0x30, 0x6a, 0x61, 0x2b, 0xff, 0xbf, 0x81,
	0x4e, 0x60, 0x80, 0x52, 0x72, 0xe9, 0xf5, 0x7d, 0x32, 0xb7, 0x93, 0x16, 0xb0, 0x4f, 0x02, 0xc3,
	0xc8, 0xb0, 0xa8, 0x42, 0xbd, 0xa1, 0x4e, 0xb4, 0x93, 0x16, 0xd0, 0x33, 0xb0, 0x54, 0x2d, 0x50,
	0xdf, 0x1d, 0x85, 0xd3, 0x7d, 0x35, 0x41, 0x5a, 0x0b, 0x4c, 0xb4, 0x8d, 0xfa, 0xe0, 0xe4, 0x58,
	0x65, 0xb2, 0x10, 0xaa, 0xe0, 0xa5, 0x77, 0xa0, 0xa3, 0x4c, 0x8a, 0x9d, 0x82, 0xd5, 0xf8, 0xa9,
	0x0d, 0x83, 0xe5, 0xea, 0xfa, 0x36, 0x76, 0x7b, 0x74, 0x08, 0x56, 0x1a, 0x3f, 0xa6, 0x2e, 0x69,
	0xc8, 0xf5, 0x32, 0x8a, 0xef, 0xdc, 0x7e, 0xf8, 0x41, 0xc0, 0x5e, 0xd5, 0xf7, 0x6d, 0x17, 0xbd,
	0x84, 0x81, 0x5e, 0x88, 0x7a, 0xe6, 0x0b, 0xcc, 0xd9, 0x67, 0xd3, 0x3d, 0x4a, 0x37, 0x67, 0x8f,
	0x5e, 0x80, 0xd5, 0x2c, 0x48, 0x4f, 0x4c, 0x93, 0x31, 0xf1, 0xcc, 0xfb, 0x2b, 0x7c, 0x1f, 0x2f,
	0xc8, 0xf3, 0xa1, 0xfe, 0xf5, 0xf3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xf0, 0x28, 0x22,
	0x09, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MyPackageClient is the client API for MyPackage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MyPackageClient interface {
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (MyPackage_FindClient, error)
}

type myPackageClient struct {
	cc *grpc.ClientConn
}

func NewMyPackageClient(cc *grpc.ClientConn) MyPackageClient {
	return &myPackageClient{cc}
}

func (c *myPackageClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/my_package.MyPackage/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myPackageClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (MyPackage_FindClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MyPackage_serviceDesc.Streams[0], "/my_package.MyPackage/Find", opts...)
	if err != nil {
		return nil, err
	}
	x := &myPackageFindClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MyPackage_FindClient interface {
	Recv() (*FindResponse, error)
	grpc.ClientStream
}

type myPackageFindClient struct {
	grpc.ClientStream
}

func (x *myPackageFindClient) Recv() (*FindResponse, error) {
	m := new(FindResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyPackageServer is the server API for MyPackage service.
type MyPackageServer interface {
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
	Find(*FindRequest, MyPackage_FindServer) error
}

// UnimplementedMyPackageServer can be embedded to have forward compatible implementations.
type UnimplementedMyPackageServer struct {
}

func (*UnimplementedMyPackageServer) Store(ctx context.Context, req *StoreRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedMyPackageServer) Find(req *FindRequest, srv MyPackage_FindServer) error {
	return status.Errorf(codes.Unimplemented, "method Find not implemented")
}

func RegisterMyPackageServer(s *grpc.Server, srv MyPackageServer) {
	s.RegisterService(&_MyPackage_serviceDesc, srv)
}

func _MyPackage_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyPackageServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/my_package.MyPackage/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyPackageServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyPackage_Find_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MyPackageServer).Find(m, &myPackageFindServer{stream})
}

type MyPackage_FindServer interface {
	Send(*FindResponse) error
	grpc.ServerStream
}

type myPackageFindServer struct {
	grpc.ServerStream
}

func (x *myPackageFindServer) Send(m *FindResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MyPackage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "my_package.MyPackage",
	HandlerType: (*MyPackageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _MyPackage_Store_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Find",
			Handler:       _MyPackage_Find_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "my_package.proto",
}
