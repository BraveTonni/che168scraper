.PHONY: up

up:
	@go mod tidy && go run ./cmd/scraper
