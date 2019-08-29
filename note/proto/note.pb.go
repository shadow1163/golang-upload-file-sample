// Code generated by protoc-gen-go. DO NOT EDIT.
// source: note.proto

package note

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Message represents a simple message sent to the Echo service.
type Message struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_640dafe07df50d4e, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Message) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type MessageArray struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MessageArray) Reset()         { *m = MessageArray{} }
func (m *MessageArray) String() string { return proto.CompactTextString(m) }
func (*MessageArray) ProtoMessage()    {}
func (*MessageArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_640dafe07df50d4e, []int{1}
}

func (m *MessageArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageArray.Unmarshal(m, b)
}
func (m *MessageArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageArray.Marshal(b, m, deterministic)
}
func (m *MessageArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageArray.Merge(m, src)
}
func (m *MessageArray) XXX_Size() int {
	return xxx_messageInfo_MessageArray.Size(m)
}
func (m *MessageArray) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageArray.DiscardUnknown(m)
}

var xxx_messageInfo_MessageArray proto.InternalMessageInfo

func (m *MessageArray) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "note.Message")
	proto.RegisterType((*MessageArray)(nil), "note.MessageArray")
}

func init() { proto.RegisterFile("note.proto", fileDescriptor_640dafe07df50d4e) }

var fileDescriptor_640dafe07df50d4e = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcb, 0x2f, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3,
	0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3,
	0xf3, 0x8a, 0x21, 0x6a, 0x94, 0x02, 0xb9, 0xd8, 0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85,
	0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84,
	0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25, 0x98, 0xc0, 0x42, 0x10, 0x8e, 0x90, 0x02,
	0x17, 0x77, 0x4a, 0x6a, 0x71, 0x72, 0x51, 0x66, 0x01, 0xc8, 0x18, 0x09, 0x66, 0xb0, 0x1c, 0xb2,
	0x90, 0x92, 0x25, 0x17, 0x0f, 0xd4, 0x48, 0xc7, 0xa2, 0xa2, 0xc4, 0x4a, 0x21, 0x4d, 0x2e, 0x8e,
	0x5c, 0x08, 0xbf, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x88, 0x57, 0x0f, 0xec, 0x4a, 0xa8,
	0xaa, 0x20, 0xb8, 0xb4, 0x51, 0x1f, 0x13, 0x17, 0xb7, 0x5f, 0x7e, 0x49, 0x6a, 0x70, 0x6a, 0x51,
	0x59, 0x66, 0x72, 0xaa, 0x90, 0x2d, 0x17, 0xb3, 0x7b, 0x6a, 0x89, 0x10, 0xaa, 0x7a, 0x29, 0x21,
	0x14, 0x2e, 0xd8, 0x12, 0x25, 0xe1, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0xf1, 0x0a, 0x71, 0xeb, 0xa7,
	0x83, 0x3c, 0x0a, 0x52, 0x50, 0x2c, 0x64, 0xc3, 0xc5, 0x12, 0x90, 0x5f, 0x8c, 0xa1, 0x1f, 0x95,
	0xab, 0x24, 0x06, 0xd6, 0x2a, 0xa0, 0x84, 0xac, 0xd5, 0x8a, 0x51, 0x4b, 0xc8, 0x81, 0x8b, 0x39,
	0xa0, 0x94, 0x90, 0x66, 0x69, 0xb0, 0x66, 0x51, 0x29, 0x01, 0x24, 0xcd, 0xfa, 0xd5, 0x99, 0x29,
	0xb5, 0x10, 0x13, 0xd8, 0x5c, 0x52, 0x73, 0x52, 0x4b, 0x52, 0x09, 0x18, 0x22, 0x01, 0x36, 0x44,
	0x48, 0x0b, 0xc3, 0x10, 0x27, 0xb6, 0x28, 0x70, 0x24, 0x26, 0xb1, 0x81, 0x63, 0xcb, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x77, 0xc2, 0x9b, 0xf8, 0xdf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NoteServiceClient is the client API for NoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NoteServiceClient interface {
	// Echo method receives a simple message and returns it.
	// The message posted as the id parameter will also be returned.
	Get(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageArray, error)
	Post(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Put(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Delete(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type noteServiceClient struct {
	cc *grpc.ClientConn
}

func NewNoteServiceClient(cc *grpc.ClientConn) NoteServiceClient {
	return &noteServiceClient{cc}
}

func (c *noteServiceClient) Get(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessageArray, error) {
	out := new(MessageArray)
	err := c.cc.Invoke(ctx, "/note.NoteService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) Post(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/note.NoteService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) Put(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/note.NoteService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) Delete(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/note.NoteService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteServiceServer is the server API for NoteService service.
type NoteServiceServer interface {
	// Echo method receives a simple message and returns it.
	// The message posted as the id parameter will also be returned.
	Get(context.Context, *Message) (*MessageArray, error)
	Post(context.Context, *Message) (*Message, error)
	Put(context.Context, *Message) (*Message, error)
	Delete(context.Context, *Message) (*Message, error)
}

// UnimplementedNoteServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNoteServiceServer struct {
}

func (*UnimplementedNoteServiceServer) Get(ctx context.Context, req *Message) (*MessageArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedNoteServiceServer) Post(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (*UnimplementedNoteServiceServer) Put(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedNoteServiceServer) Delete(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterNoteServiceServer(s *grpc.Server, srv NoteServiceServer) {
	s.RegisterService(&_NoteService_serviceDesc, srv)
}

func _NoteService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Get(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Post(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Put(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.NoteService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).Delete(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _NoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "note.NoteService",
	HandlerType: (*NoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _NoteService_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _NoteService_Post_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _NoteService_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NoteService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "note.proto",
}