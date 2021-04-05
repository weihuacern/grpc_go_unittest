package server

import (
	"context"

	pb "github.com/weihuacern/grpc_test_proto/auth_user"
)

type GRPCMethod struct {
	pb.UnimplementedAuthUserServer
}

// CreateAuthUser : RPC method to create user
func (as *GRPCMethod) CreateAuthUser(ctx context.Context, in *pb.CreateAuthUserRequest) (*pb.CreateAuthUserReply, error) {
	// TODO, do someting
	return &pb.CreateAuthUserReply{Reply: &pb.AuthUserReply{Msg: "abc"}}, nil
}

// ReadAuthUser : RPC method to read user
func (as *GRPCMethod) ReadAuthUser(ctx context.Context, in *pb.ReadAuthUserRequest) (*pb.ReadAuthUserReply, error) {
	// TODO, do someting
	return &pb.ReadAuthUserReply{Reply: &pb.AuthUserReply{Msg: "abc"}}, nil
}

// UpdateAuthUser : RPC method to update user
func (as *GRPCMethod) UpdateAuthUser(ctx context.Context, in *pb.UpdateAuthUserRequest) (*pb.UpdateAuthUserReply, error) {
	// TODO, do someting
	return &pb.UpdateAuthUserReply{Reply: &pb.AuthUserReply{Msg: "abc"}}, nil
}

// DeleteAuthUser : RPC method to delete user
func (as *GRPCMethod) DeleteAuthUser(ctx context.Context, in *pb.DeleteAuthUserRequest) (*pb.DeleteAuthUserReply, error) {
	// TODO, do someting
	return &pb.DeleteAuthUserReply{Reply: &pb.AuthUserReply{Msg: "abc"}}, nil
}
