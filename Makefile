# Path to proto directories
PROTO_DIR := proto
OUT_DIR := .

# Find all .proto files in proto/ recursively
PROTO_FILES := $(shell find $(PROTO_DIR) -name "*.proto")

# Resolve path to googleapis (needed for grpc-gateway annotations)
GOOGLEAPIS := $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway/v2)/third_party/googleapis

generate:
	@echo "Generating Go, gRPC, and Gateway code..."
	protoc -I $(PROTO_DIR) -I $(GOOGLEAPIS) \
	       --go_out $(OUT_DIR) --go_opt paths=source_relative \
	       --go-grpc_out $(OUT_DIR) --go-grpc_opt paths=source_relative \
	       --grpc-gateway_out $(OUT_DIR) --grpc-gateway_opt paths=source_relative \
	       $(PROTO_FILES)

clean:
	@echo "Cleaning generated .pb.go files..."
	find $(PROTO_DIR) -name "*.pb.go" -delete

.PHONY: generate clean
