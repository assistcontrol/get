#!/bin/sh

title() { echo "========> $@"; }

title go test
go test ./...

title go vet
go vet ./...

title golangci-lint
golangci-lint run ./...