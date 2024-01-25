protogen:
	@./scripts/protogen.sh

launch:
	@make clear
	@go run main.go partition 8000 test

launch2:
	@make clear
	@go run main.go partition 7999 test2

blaunch:
	@make clear
	@go run main.go balancer 8001 1

register-partition:
	@echo '{"address": "127.0.0.1:8000"}' | grpcurl -d @ -v -plaintext localhost:8001 dkvs.balancer.BalancerService/RegisterPartition
	@echo '{"address": "127.0.0.1:7999"}' | grpcurl -d @ -v -plaintext localhost:8001 dkvs.balancer.BalancerService/RegisterPartition

setb:
	@echo '{"key": "a2V5", "value": "ZGF0YQ==", "lamport": ${L}}' | grpcurl -d @ -v -plaintext localhost:8001 dkvs.balancer.BalancerService/Set

getb:
	@echo '{"key": "a2V5", "lamport": ${L}}' | grpcurl -d @ -v -plaintext localhost:8001 dkvs.balancer.BalancerService/Get

shr:
	@echo '{"min": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==", "max": "/////////////////////w=="}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/SetHashrange

set:
	@echo '{"key": "a2V5", "value": "ZGF0YQ==", "lamport": ${L}}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/Set

get:
	@echo '{"key": "a2V5", "lamport": ${L}}' | grpcurl -d @ -v -plaintext localhost:8000 dkvs.partition.PartitionService/Get

clear:
	@rm -rf test
	@rm -rf test2
	@rm -rf balancer-db

unit:
	@go test -v ./...