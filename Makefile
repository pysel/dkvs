protogen:
	@./scripts/protogen.sh

# echo '{"min": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", "max": "/////////////////////w=="}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/SetHashrange
# echo '{"key": "a2V5", "value": "ZGF0YQ==", "lamport": 1}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/Set
# echo '{"key": "a2V5"}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/Get
launch:
	@go run main.go partition 8000 test

shr:
	@echo '{"min": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", "max": "/////////////////////w=="}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/SetHashrange

set:
	@echo '{"key": "a2V5", "value": "ZGF0YQ==", "lamport": ${L}}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/Set

get:
	@echo '{"key": "a2V5", "lamport": ${L}}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/Get

unit:
	@go test -v ./...