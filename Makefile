GOBINARYNAME := grpc_go_unittest

all: rpc_go

## rpc_go: Build Go gRPC server
rpc_go:
	golint
	go build -o $(GOBINARYNAME)

## clean: Clean up
clean:
	rm -rf $(GOBINARYNAME)

## help: Obtain help related information
help: Makefile
	@sed -n 's/^##//p' $<
