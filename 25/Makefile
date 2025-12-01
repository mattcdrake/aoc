.DEFAULT_GOAL := build

.PHONY: clean fmt vet run build
clean:
	rm -f aoc25
	go clean
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
run: clean vet
	go run main.go d1.go
build: clean vet
	go build -o aoc25
