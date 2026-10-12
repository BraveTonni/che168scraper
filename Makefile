.PHONY: up test

up:
	@go mod tidy && go run ./cmd/scraper

test:
	@go mod tidy && go test -v ./...
