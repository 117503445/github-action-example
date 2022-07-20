# github-action-example

This repo demonstrates the ability of github action to do docker image builds and github releases using a simple go http code.

<https://github.com/117503445/github-action-example/blob/master/.github/workflows/docker.yaml> is a docker build template that can be easily copied to other projects.

[![codecov](https://codecov.io/gh/117503445/github-action-example/branch/master/graph/badge.svg)](https://codecov.io/gh/117503445/github-action-example)

## dev

```sh
go run .
go test ./... -race -coverpkg=./... -coverprofile=coverage.txt -v
```
