# Demo CI Pipelines
Multiple CI pipelines for a demo application in Go.

| CI Tool        | Status                                                                                                                                                                                         | Configuration                                                                                               |
|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
| GitHub Actions | [![GitHub Actions Demo](https://github.com/PhilomathesInc/demo-ci-pipelines/actions/workflows/github-actions-demo.yml/badge.svg)](https://github.com/PhilomathesInc/demo-ci-pipelines/actions) | [Demo workflow](./github/workflows/github-actions-demo.yaml) that builds the Go application server and CLI on GitHub-hosted Ubuntu Linux runner  |
| Travis CI      | [![Build Status](https://app.travis-ci.com/PhilomathesInc/demo-ci-pipelines.svg?branch=main)](https://app.travis-ci.com/PhilomathesInc/demo-ci-pipelines)                                      | [Demo build config](./travis.yml) that builds the Go application server and CLI in Ubuntu Linux environment        |
| CircleCI       | [![CircleCI](https://circleci.com/gh/PhilomathesInc/demo-ci-pipelines.svg?style=svg)](https://app.circleci.com/pipelines/github/PhilomathesInc/demo-ci-pipelines)                              | [Demo build config](./circleci/config.yml) that builds the Go application server and CLI on docker executor |

### Go Application
A demo application is written in Go that has a server and a CLI.
- For creating the server, [net/http](https://golang.org/pkg/net/http) standard package is used.
- For creating the CLI, [Cobra](https://github.com/spf13/cobra) library is used.
- For logging, [Zap](https://github.com/uber-go/zap) library is used.

#### Building and running the server
```
make build
make run
```

#### Building and running the CLI
```
make cli
./demo-cli --help
```
