build:
	@go build -o bin/casgo

run: build
	@./bin/casgo
test:
	@go test -v -json ./... | tparse -all

.PHONY: build run
