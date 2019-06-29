// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rustcompile.proto

package rustcompile

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Code struct {
	Files                []*File  `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Code) Reset()         { *m = Code{} }
func (m *Code) String() string { return proto.CompactTextString(m) }
func (*Code) ProtoMessage()    {}
func (*Code) Descriptor() ([]byte, []int) {
	return fileDescriptor_194bd31a0c7a6409, []int{0}
}

func (m *Code) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Code.Unmarshal(m, b)
}
func (m *Code) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Code.Marshal(b, m, deterministic)
}
func (m *Code) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Code.Merge(m, src)
}
func (m *Code) XXX_Size() int {
	return xxx_messageInfo_Code.Size(m)
}
func (m *Code) XXX_DiscardUnknown() {
	xxx_messageInfo_Code.DiscardUnknown(m)
}

var xxx_messageInfo_Code proto.InternalMessageInfo

func (m *Code) GetFiles() []*File {
	if m != nil {
		return m.Files
	}
	return nil
}

type File struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_194bd31a0c7a6409, []int{1}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *File) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Result struct {
	Log                  string   `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
	Binary               []byte   `protobuf:"bytes,2,opt,name=binary,proto3" json:"binary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_194bd31a0c7a6409, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *Result) GetBinary() []byte {
	if m != nil {
		return m.Binary
	}
	return nil
}

func init() {
	proto.RegisterType((*Code)(nil), "rustcompile.Code")
	proto.RegisterType((*File)(nil), "rustcompile.File")
	proto.RegisterType((*Result)(nil), "rustcompile.Result")
}

func init() { proto.RegisterFile("rustcompile.proto", fileDescriptor_194bd31a0c7a6409) }

var fileDescriptor_194bd31a0c7a6409 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x3f, 0x8f, 0x82, 0x40,
	0x14, 0xc4, 0x6f, 0x0f, 0x8e, 0xe4, 0x1e, 0x57, 0x1c, 0xcf, 0xc4, 0x10, 0x2b, 0xb2, 0x8d, 0x54,
	0x68, 0xb0, 0xb1, 0x96, 0x44, 0xfb, 0xf5, 0x13, 0x80, 0xac, 0xba, 0xc9, 0xea, 0x92, 0xfd, 0x53,
	0xf8, 0xed, 0xcd, 0x02, 0x26, 0xd0, 0xcd, 0xcc, 0xcb, 0x2f, 0x33, 0x0f, 0x12, 0xed, 0x8c, 0xbd,
	0xa8, 0x47, 0x27, 0x24, 0x2f, 0x3a, 0xad, 0xac, 0xc2, 0x78, 0x12, 0xd1, 0x0d, 0x84, 0x95, 0x6a,
	0x39, 0xae, 0xe1, 0xe7, 0x2a, 0x24, 0x37, 0x29, 0xc9, 0x82, 0x3c, 0x2e, 0x93, 0x62, 0xca, 0x1d,
	0x85, 0xe4, 0x6c, 0xb8, 0xd3, 0x02, 0x42, 0x6f, 0x11, 0x21, 0xec, 0x6a, 0x7b, 0x4f, 0x49, 0x46,
	0xf2, 0x5f, 0xd6, 0x6b, 0x9f, 0x35, 0xaa, 0x7d, 0xa5, 0xdf, 0x19, 0xc9, 0xff, 0x58, 0xaf, 0x69,
	0x09, 0x11, 0xe3, 0xc6, 0x49, 0x8b, 0xff, 0x10, 0x48, 0x75, 0x1b, 0x01, 0x2f, 0x71, 0x09, 0x51,
	0x23, 0x9e, 0xb5, 0xfe, 0x10, 0xa3, 0x2b, 0x4f, 0x10, 0x33, 0x67, 0x6c, 0x35, 0xd4, 0xe3, 0x1e,
	0xe0, 0x6c, 0x6b, 0x6d, 0x0f, 0x4e, 0xc8, 0x16, 0xe7, 0xd3, 0xfc, 0xf8, 0xd5, 0x62, 0x16, 0x0d,
	0x75, 0xf4, 0x6b, 0x4b, 0x9a, 0xa8, 0xff, 0x78, 0xf7, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xfe,
	0xd1, 0x55, 0x06, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RustCompileClient is the client API for RustCompile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RustCompileClient interface {
	StartBuild(ctx context.Context, in *Code, opts ...grpc.CallOption) (RustCompile_StartBuildClient, error)
}

type rustCompileClient struct {
	cc *grpc.ClientConn
}

func NewRustCompileClient(cc *grpc.ClientConn) RustCompileClient {
	return &rustCompileClient{cc}
}

func (c *rustCompileClient) StartBuild(ctx context.Context, in *Code, opts ...grpc.CallOption) (RustCompile_StartBuildClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RustCompile_serviceDesc.Streams[0], "/rustcompile.RustCompile/StartBuild", opts...)
	if err != nil {
		return nil, err
	}
	x := &rustCompileStartBuildClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RustCompile_StartBuildClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type rustCompileStartBuildClient struct {
	grpc.ClientStream
}

func (x *rustCompileStartBuildClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RustCompileServer is the server API for RustCompile service.
type RustCompileServer interface {
	StartBuild(*Code, RustCompile_StartBuildServer) error
}

// UnimplementedRustCompileServer can be embedded to have forward compatible implementations.
type UnimplementedRustCompileServer struct {
}

func (*UnimplementedRustCompileServer) StartBuild(req *Code, srv RustCompile_StartBuildServer) error {
	return status.Errorf(codes.Unimplemented, "method StartBuild not implemented")
}

func RegisterRustCompileServer(s *grpc.Server, srv RustCompileServer) {
	s.RegisterService(&_RustCompile_serviceDesc, srv)
}

func _RustCompile_StartBuild_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Code)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RustCompileServer).StartBuild(m, &rustCompileStartBuildServer{stream})
}

type RustCompile_StartBuildServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type rustCompileStartBuildServer struct {
	grpc.ServerStream
}

func (x *rustCompileStartBuildServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

var _RustCompile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rustcompile.RustCompile",
	HandlerType: (*RustCompileServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartBuild",
			Handler:       _RustCompile_StartBuild_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rustcompile.proto",
}
