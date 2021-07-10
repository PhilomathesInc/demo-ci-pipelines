.PHONY: cli start

cli:
	go build -o demo-cli cli/main.go

start:
	go run main.go
