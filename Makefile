protogen:
	@./scripts/protogen.sh

launch:
	@make clear
	@go run main.go partition 8000 test

shr:
	@echo '{"min": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", "max": "/////////////////////w=="}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/SetHashrange

set:
	@echo '{"key": "a2V5", "value": "ZGF0YQ==", "lamport": ${L}}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/Set

get:
	@echo '{"key": "a2V5", "lamport": ${L}}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/Get

clear:
	@rm -rf test

unit:
	@go test -v ./...