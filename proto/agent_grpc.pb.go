// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v6.31.1
// source: proto/agent.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	SendConfig(ctx context.Context, in *AgentConfigRequest, opts ...grpc.CallOption) (*AgentCommandResponse, error)
	SendResult(ctx context.Context, in *AgentCommandResultRequest, opts ...grpc.CallOption) (*AcknowledgementResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) SendConfig(ctx context.Context, in *AgentConfigRequest, opts ...grpc.CallOption) (*AgentCommandResponse, error) {
	out := new(AgentCommandResponse)
	err := c.cc.Invoke(ctx, "/agent.AgentService/SendConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) SendResult(ctx context.Context, in *AgentCommandResultRequest, opts ...grpc.CallOption) (*AcknowledgementResponse, error) {
	out := new(AcknowledgementResponse)
	err := c.cc.Invoke(ctx, "/agent.AgentService/SendResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	SendConfig(context.Context, *AgentConfigRequest) (*AgentCommandResponse, error)
	SendResult(context.Context, *AgentCommandResultRequest) (*AcknowledgementResponse, error)
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) SendConfig(context.Context, *AgentConfigRequest) (*AgentCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendConfig not implemented")
}
func (UnimplementedAgentServiceServer) SendResult(context.Context, *AgentCommandResultRequest) (*AcknowledgementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendResult not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_SendConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).SendConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentService/SendConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).SendConfig(ctx, req.(*AgentConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_SendResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCommandResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).SendResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.AgentService/SendResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).SendResult(ctx, req.(*AgentCommandResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendConfig",
			Handler:    _AgentService_SendConfig_Handler,
		},
		{
			MethodName: "SendResult",
			Handler:    _AgentService_SendResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/agent.proto",
}
