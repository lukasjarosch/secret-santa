.PHONY: run
run:
	@go run cmd/secretsanta/main.go

.PHONY: proto
proto:
	@echo "building protobuf stubs into: ./api"
	@cd proto; ./build.sh

.PHONY: grpcui
grpcui:
	@cd proto; grpcui -import-path . -import-path ./third_party -proto ./service.proto -port 5000 -plaintext localhost:50051
