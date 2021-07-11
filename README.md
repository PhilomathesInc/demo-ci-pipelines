# Demo CI Pipelines
Multiple CI pipelines for a demo application in Go.

[![CircleCI](https://circleci.com/gh/PhilomathesInc/demo-ci-pipelines.svg?style=svg)](https://circleci.com/gh/circleci/circleci-docs)

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

## GitHub Actions
[Demo workflow](./github/workflows/github-actions-demo.yaml) that builds the Go application server and CLI.

## Travis CI
[Demo build config](./travis.yml) that builds the Go application server and CLI on linux and OS X.

**Note:** Always run your Travis CI configuration through the [Travis CI Build Config Explorer](https://config.travis-ci.com/explore)

## CircleCI
[Demo build config](./circleci/config.yml) that builds the Go application server and CLI on docker executor.
