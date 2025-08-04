# echo-mvc-generator

[![CI/CD](https://github.com/sk1t0n/echo-mvc-generator/actions/workflows/go.yml/badge.svg)](https://github.com/sk1t0n/echo-mvc-generator/actions/workflows/go.yml)

CLI for creating a model, view or controller. The paths are the same as in the template [echo-mvc-template](https://github.com/sk1t0n/echo-mvc-template). If you wish, you can change the paths to your own. Tech stack: [Cobra](https://github.com/spf13/cobra).

## Install

```sh
go install github.com/sk1t0n/echo-mvc-generator@latest
```

## Run CLI

```sh
# you need to add ~/go/bin to your PATH environment variable
echo-mvc-generator
```

## Run tests

```sh
make run_tests
```

## Show code coverage

```sh
make run_cover_tests
```
