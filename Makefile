.PHONY: build
build:
	go build -v ./cmd/vk_photo_downloader

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build