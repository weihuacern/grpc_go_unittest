module grpc_go_unittest

go 1.13

require (
	github.com/agiledragon/gomonkey/v2 v2.1.0
	github.com/golang/protobuf v1.4.2
	github.com/stretchr/testify v1.6.1
	github.com/weihuacern/grpc_test_proto/auth_user v1.0.1
	google.golang.org/grpc v1.36.1
)

require (
	client v0.0.0
	server v0.0.0
)

replace (
	client => ./src/client
	server => ./src/server
)
