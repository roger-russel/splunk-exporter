.PHONY: run
run:
	@go run cmd/splunk-exporter/main.go

.PHONY: test
test:
	@go test ./... -race

.PHONY: lint
lint:
	@golangci-lint run