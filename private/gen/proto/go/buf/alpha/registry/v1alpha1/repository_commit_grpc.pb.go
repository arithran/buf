// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             (unknown)
// source: buf/alpha/registry/v1alpha1/repository_commit.proto

package registryv1alpha1

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

// RepositoryCommitServiceClient is the client API for RepositoryCommitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepositoryCommitServiceClient interface {
	// ListRepositoryCommitsByBranch lists the repository commits associated
	// with a repository branch on a repository, ordered by their create time.
	ListRepositoryCommitsByBranch(ctx context.Context, in *ListRepositoryCommitsByBranchRequest, opts ...grpc.CallOption) (*ListRepositoryCommitsByBranchResponse, error)
	// ListRepositoryCommitsByReference returns repository commits up-to and including
	// the provided reference.
	ListRepositoryCommitsByReference(ctx context.Context, in *ListRepositoryCommitsByReferenceRequest, opts ...grpc.CallOption) (*ListRepositoryCommitsByReferenceResponse, error)
	// GetRepositoryCommitByReference returns the repository commit matching
	// the provided reference, if it exists.
	GetRepositoryCommitByReference(ctx context.Context, in *GetRepositoryCommitByReferenceRequest, opts ...grpc.CallOption) (*GetRepositoryCommitByReferenceResponse, error)
	// GetRepositoryCommitBySequenceId returns the repository commit matching
	// the provided sequence ID and branch, if it exists.
	GetRepositoryCommitBySequenceId(ctx context.Context, in *GetRepositoryCommitBySequenceIdRequest, opts ...grpc.CallOption) (*GetRepositoryCommitBySequenceIdResponse, error)
}

type repositoryCommitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryCommitServiceClient(cc grpc.ClientConnInterface) RepositoryCommitServiceClient {
	return &repositoryCommitServiceClient{cc}
}

func (c *repositoryCommitServiceClient) ListRepositoryCommitsByBranch(ctx context.Context, in *ListRepositoryCommitsByBranchRequest, opts ...grpc.CallOption) (*ListRepositoryCommitsByBranchResponse, error) {
	out := new(ListRepositoryCommitsByBranchResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.RepositoryCommitService/ListRepositoryCommitsByBranch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryCommitServiceClient) ListRepositoryCommitsByReference(ctx context.Context, in *ListRepositoryCommitsByReferenceRequest, opts ...grpc.CallOption) (*ListRepositoryCommitsByReferenceResponse, error) {
	out := new(ListRepositoryCommitsByReferenceResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.RepositoryCommitService/ListRepositoryCommitsByReference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryCommitServiceClient) GetRepositoryCommitByReference(ctx context.Context, in *GetRepositoryCommitByReferenceRequest, opts ...grpc.CallOption) (*GetRepositoryCommitByReferenceResponse, error) {
	out := new(GetRepositoryCommitByReferenceResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.RepositoryCommitService/GetRepositoryCommitByReference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryCommitServiceClient) GetRepositoryCommitBySequenceId(ctx context.Context, in *GetRepositoryCommitBySequenceIdRequest, opts ...grpc.CallOption) (*GetRepositoryCommitBySequenceIdResponse, error) {
	out := new(GetRepositoryCommitBySequenceIdResponse)
	err := c.cc.Invoke(ctx, "/buf.alpha.registry.v1alpha1.RepositoryCommitService/GetRepositoryCommitBySequenceId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryCommitServiceServer is the server API for RepositoryCommitService service.
// All implementations should embed UnimplementedRepositoryCommitServiceServer
// for forward compatibility
type RepositoryCommitServiceServer interface {
	// ListRepositoryCommitsByBranch lists the repository commits associated
	// with a repository branch on a repository, ordered by their create time.
	ListRepositoryCommitsByBranch(context.Context, *ListRepositoryCommitsByBranchRequest) (*ListRepositoryCommitsByBranchResponse, error)
	// ListRepositoryCommitsByReference returns repository commits up-to and including
	// the provided reference.
	ListRepositoryCommitsByReference(context.Context, *ListRepositoryCommitsByReferenceRequest) (*ListRepositoryCommitsByReferenceResponse, error)
	// GetRepositoryCommitByReference returns the repository commit matching
	// the provided reference, if it exists.
	GetRepositoryCommitByReference(context.Context, *GetRepositoryCommitByReferenceRequest) (*GetRepositoryCommitByReferenceResponse, error)
	// GetRepositoryCommitBySequenceId returns the repository commit matching
	// the provided sequence ID and branch, if it exists.
	GetRepositoryCommitBySequenceId(context.Context, *GetRepositoryCommitBySequenceIdRequest) (*GetRepositoryCommitBySequenceIdResponse, error)
}

// UnimplementedRepositoryCommitServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRepositoryCommitServiceServer struct {
}

func (UnimplementedRepositoryCommitServiceServer) ListRepositoryCommitsByBranch(context.Context, *ListRepositoryCommitsByBranchRequest) (*ListRepositoryCommitsByBranchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepositoryCommitsByBranch not implemented")
}
func (UnimplementedRepositoryCommitServiceServer) ListRepositoryCommitsByReference(context.Context, *ListRepositoryCommitsByReferenceRequest) (*ListRepositoryCommitsByReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepositoryCommitsByReference not implemented")
}
func (UnimplementedRepositoryCommitServiceServer) GetRepositoryCommitByReference(context.Context, *GetRepositoryCommitByReferenceRequest) (*GetRepositoryCommitByReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepositoryCommitByReference not implemented")
}
func (UnimplementedRepositoryCommitServiceServer) GetRepositoryCommitBySequenceId(context.Context, *GetRepositoryCommitBySequenceIdRequest) (*GetRepositoryCommitBySequenceIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepositoryCommitBySequenceId not implemented")
}

// UnsafeRepositoryCommitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepositoryCommitServiceServer will
// result in compilation errors.
type UnsafeRepositoryCommitServiceServer interface {
	mustEmbedUnimplementedRepositoryCommitServiceServer()
}

func RegisterRepositoryCommitServiceServer(s grpc.ServiceRegistrar, srv RepositoryCommitServiceServer) {
	s.RegisterService(&RepositoryCommitService_ServiceDesc, srv)
}

func _RepositoryCommitService_ListRepositoryCommitsByBranch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepositoryCommitsByBranchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryCommitServiceServer).ListRepositoryCommitsByBranch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.RepositoryCommitService/ListRepositoryCommitsByBranch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryCommitServiceServer).ListRepositoryCommitsByBranch(ctx, req.(*ListRepositoryCommitsByBranchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryCommitService_ListRepositoryCommitsByReference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepositoryCommitsByReferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryCommitServiceServer).ListRepositoryCommitsByReference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.RepositoryCommitService/ListRepositoryCommitsByReference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryCommitServiceServer).ListRepositoryCommitsByReference(ctx, req.(*ListRepositoryCommitsByReferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryCommitService_GetRepositoryCommitByReference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepositoryCommitByReferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryCommitServiceServer).GetRepositoryCommitByReference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.RepositoryCommitService/GetRepositoryCommitByReference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryCommitServiceServer).GetRepositoryCommitByReference(ctx, req.(*GetRepositoryCommitByReferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepositoryCommitService_GetRepositoryCommitBySequenceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepositoryCommitBySequenceIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryCommitServiceServer).GetRepositoryCommitBySequenceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buf.alpha.registry.v1alpha1.RepositoryCommitService/GetRepositoryCommitBySequenceId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryCommitServiceServer).GetRepositoryCommitBySequenceId(ctx, req.(*GetRepositoryCommitBySequenceIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RepositoryCommitService_ServiceDesc is the grpc.ServiceDesc for RepositoryCommitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RepositoryCommitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buf.alpha.registry.v1alpha1.RepositoryCommitService",
	HandlerType: (*RepositoryCommitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRepositoryCommitsByBranch",
			Handler:    _RepositoryCommitService_ListRepositoryCommitsByBranch_Handler,
		},
		{
			MethodName: "ListRepositoryCommitsByReference",
			Handler:    _RepositoryCommitService_ListRepositoryCommitsByReference_Handler,
		},
		{
			MethodName: "GetRepositoryCommitByReference",
			Handler:    _RepositoryCommitService_GetRepositoryCommitByReference_Handler,
		},
		{
			MethodName: "GetRepositoryCommitBySequenceId",
			Handler:    _RepositoryCommitService_GetRepositoryCommitBySequenceId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buf/alpha/registry/v1alpha1/repository_commit.proto",
}
