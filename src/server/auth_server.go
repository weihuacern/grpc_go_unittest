package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "github.com/weihuacern/grpc_test_proto/auth_user"
)

const (
	port = 2021
)

// AuthServer : gRPC server for authentication and authorization
type AuthServer struct {
	server *grpc.Server
	method *GRPCMethod
}

// Init : Initialize gRPC server
func (as *AuthServer) Init() error {
	as.server = grpc.NewServer()
	as.method = &GRPCMethod{}
	pb.RegisterAuthUserServer(as.server, as.method)
	return nil
}

// Work : Run gRPC server
func (as *AuthServer) Work() error {
	var listen net.Listener
	var err error
	if listen, err = net.Listen("tcp", fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("Failed to listen: %s\n", err.Error())
		return err
	}

	if err = as.server.Serve(listen); err != nil {
		fmt.Printf("Failed to serve: %s\n", err.Error())
		return err
	}
	fmt.Printf("Server started at Port[%d]\n", port)
	return nil
}

// NewAuthServer : Entrypoint to create new gRPC service
func NewAuthServer() (*AuthServer, error) {
	var as AuthServer
	return &as, nil
}
