# This Project contain this package

## Golang

go version -> go1.24.1

### Testify

go get github.com/stretchr/testify

### Assert

go get github.com/stretchr/testify/assert

### Mock

go get github.com/stretchr/testify/mock

### GRPC

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

export PATH="$PATH:$(go env GOPATH)/bin"

## Gorm

go get gorm.io/gorm
go get gorm.io/driver/postgres

## PostgreSQL

docker image -> postgres:17-alpine

# proto directories in this repo is only for example purpose
