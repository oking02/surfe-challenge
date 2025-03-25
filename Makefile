

setup: setup
	go get -tool -modfile=go.tool.mod github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.1

.PHONY: fmt
fmt:
	gofmt -w -s .
	goimports -w .
	go clean ./...

lint:
	go tool -modfile=go.tool.mod golangci-lint run
	@echo "Linter Passed"

.PHONY: run
run: fmt
	go run cmd/main.go

.env:
	cp .env.dist .env