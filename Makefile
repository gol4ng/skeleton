proto:
	@echo ">> Generate proto MyPackage"
	docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd) protobuf-go \
		protoc \
			-I ./pkg/my_package/protos my_package.proto \
			--go_out=plugins=grpc:./pkg/my_package/protos
