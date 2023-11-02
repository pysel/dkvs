protogen:
	@python3 scripts/protogen.py
	@go mod tidy

# echo '{"key": "key", "value": "c29tZV9kYXRhX2hlcmU="}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/SetMessage 
# echo '{"key": "key"}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/GetMessage
launch:
	@go run main.go partition 8000 test 0 50 