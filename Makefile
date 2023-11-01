protogen:
	@python3 scripts/protogen.py
	@go mod tidy

launch:
	@go run main.go partition 8000 test 0 50 