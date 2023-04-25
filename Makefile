.PHONY: launch
run:
	go run -v ./cmd/server/main.go

.PHONY: build
build:
	go build -v ./main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := run


