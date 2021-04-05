package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/weihuacern/grpc_test_proto/auth_user"
)

func TestGRPCMethod_CreateAuthUser(t *testing.T) {
	gm := &GRPCMethod{}
	ctx := context.Background()
	reply, err := gm.CreateAuthUser(
		ctx,
		&pb.CreateAuthUserRequest{})
	assert.Equal(t, reply.GetReply().GetMsg(), "abc")
	assert.Nil(t, err)
}
