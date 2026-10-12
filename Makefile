.PHONY: up test test-unit test-integration

up:
	@go mod tidy && go run ./cmd/scraper

test-unit:
	@go test -v ./...

test-integration:
	@go test -v -tags=integration ./...

test: test-unit test-integration