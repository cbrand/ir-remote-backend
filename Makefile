
protobuf:
	mkdir -p protocol
	protoc definition.proto --go_out=. --go-grpc_out=.

