BIN_DIR := build
BIN     := chescraper

.PHONY: build build-linux build-windows tidy clean

build: build-linux build-windows

build-linux:
	@set GOOS=linux&& go build -o $(BIN_DIR)/$(BIN) ./cmd/scraper

build-windows:
	@set GOOS=windows&& go build -o $(BIN_DIR)/$(BIN).exe ./cmd/scraper

up:
	@go mod tidy && go run ./cmd/scraper

test:
	@go mod tidy && go test -v ./...
