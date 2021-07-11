# Demo CI Pipelines
Multiple CI pipelines for a demo application in Go.

## Go Application
A demo application is written in Go that has a server and a CLI.
- For creating the server, [net/http](https://golang.org/pkg/net/http) standard package is used.
- For creating the CLI, [Cobra](https://github.com/spf13/cobra) library is used.
- For logging, [Zap](https://github.com/uber-go/zap) library is used.

### Building and running the server
```
make build
make run
```

### Building and running the CLI
```
make cli
./demo-cli --help
```
