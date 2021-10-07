// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: api/proto/v1/chat.proto

package __

import (
	context "context"
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

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SendersName string `protobuf:"bytes,2,opt,name=senders_name,json=sendersName,proto3" json:"senders_name,omitempty"`
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Channel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Channel) GetSendersName() string {
	if x != nil {
		return x.SendersName
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender  string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Channel *Channel `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_chat_proto_msgTypes[1]
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
	return file_api_proto_v1_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Message) GetChannel() *Channel {
	if x != nil {
		return x.Channel
	}
	return nil
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type MessageAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *MessageAck) Reset() {
	*x = MessageAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAck) ProtoMessage() {}

func (x *MessageAck) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAck.ProtoReflect.Descriptor instead.
func (*MessageAck) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_chat_proto_rawDescGZIP(), []int{2}
}

func (x *MessageAck) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_api_proto_v1_chat_proto protoreflect.FileDescriptor

var file_api_proto_v1_chat_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x40, 0x0a,
	0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x62, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63,
	0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x6a, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2e, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63,
	0x6b, 0x22, 0x00, 0x28, 0x01, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_chat_proto_rawDescOnce sync.Once
	file_api_proto_v1_chat_proto_rawDescData = file_api_proto_v1_chat_proto_rawDesc
)

func file_api_proto_v1_chat_proto_rawDescGZIP() []byte {
	file_api_proto_v1_chat_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_chat_proto_rawDescData)
	})
	return file_api_proto_v1_chat_proto_rawDescData
}

var file_api_proto_v1_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_v1_chat_proto_goTypes = []interface{}{
	(*Channel)(nil),    // 0: v1.Channel
	(*Message)(nil),    // 1: v1.Message
	(*MessageAck)(nil), // 2: v1.MessageAck
}
var file_api_proto_v1_chat_proto_depIdxs = []int32{
	0, // 0: v1.Message.channel:type_name -> v1.Channel
	0, // 1: v1.ChatService.JoinChannel:input_type -> v1.Channel
	1, // 2: v1.ChatService.SendMessage:input_type -> v1.Message
	1, // 3: v1.ChatService.JoinChannel:output_type -> v1.Message
	2, // 4: v1.ChatService.SendMessage:output_type -> v1.MessageAck
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_v1_chat_proto_init() }
func file_api_proto_v1_chat_proto_init() {
	if File_api_proto_v1_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
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
		file_api_proto_v1_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_proto_v1_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAck); i {
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
			RawDescriptor: file_api_proto_v1_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_v1_chat_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_chat_proto_depIdxs,
		MessageInfos:      file_api_proto_v1_chat_proto_msgTypes,
	}.Build()
	File_api_proto_v1_chat_proto = out.File
	file_api_proto_v1_chat_proto_rawDesc = nil
	file_api_proto_v1_chat_proto_goTypes = nil
	file_api_proto_v1_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	JoinChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChatService_JoinChannelClient, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) JoinChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChatService_JoinChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/v1.ChatService/JoinChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceJoinChannelClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_JoinChannelClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceJoinChannelClient struct {
	grpc.ClientStream
}

func (x *chatServiceJoinChannelClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[1], "/v1.ChatService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSendMessageClient{stream}
	return x, nil
}

type ChatService_SendMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*MessageAck, error)
	grpc.ClientStream
}

type chatServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceSendMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSendMessageClient) CloseAndRecv() (*MessageAck, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	JoinChannel(*Channel, ChatService_JoinChannelServer) error
	SendMessage(ChatService_SendMessageServer) error
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) JoinChannel(*Channel, ChatService_JoinChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinChannel not implemented")
}
func (*UnimplementedChatServiceServer) SendMessage(ChatService_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_JoinChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Channel)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).JoinChannel(m, &chatServiceJoinChannelServer{stream})
}

type ChatService_JoinChannelServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatServiceJoinChannelServer struct {
	grpc.ServerStream
}

func (x *chatServiceJoinChannelServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).SendMessage(&chatServiceSendMessageServer{stream})
}

type ChatService_SendMessageServer interface {
	SendAndClose(*MessageAck) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceSendMessageServer) SendAndClose(m *MessageAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSendMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinChannel",
			Handler:       _ChatService_JoinChannel_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMessage",
			Handler:       _ChatService_SendMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/proto/v1/chat.proto",
}
