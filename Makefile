.PHONY: build start cli

build:
	go build -o demo-server main.go

start: build
	./demo-server

cli:
	go build -o demo-cli cli/main.go
