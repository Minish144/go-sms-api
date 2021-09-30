generate:
	rm -rf gen/{pb,clean_proto,swagger}/{*.go,*.proto,*.json}
	mkdir -p gen/{pb,swagger}

	protoc \
		-I=proto/ \
		-I=${GOPATH}/src/ \
		-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=./ \
		--validate_out=lang=go:./ \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=allow_merge=true,merge_file_name=api:./gen/swagger/ proto/*.proto

	protoc \
		-I=proto/ \
		-I=${GOPATH}/src/ \
		-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go-grpc_out=./ proto/api.proto
