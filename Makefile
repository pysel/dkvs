protogen:
	@./scripts/protogen.sh

# echo '{"key": "key", "value": "c29tZV9kYXRhX2hlcmU="}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/SetMessage 
# echo '{"key": "key"}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/GetMessage
launch:
	@go run main.go partition 8000 test

unit:
	@go test -v ./...