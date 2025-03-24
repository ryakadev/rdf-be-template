FROM golang:1.24.1-alpine AS build

ENV CGO_ENABLED=1
ENV GOOS=linux
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk add --no-cache \
    git \
    bash \
    build-base \
    protobuf-dev \ 
    protobuf \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# RUN protoc -I ./proto ./proto/*.proto --go_out=./ --go-grpc_out=./

RUN protoc --go_out=. --go-grpc_out=. proto/*.proto

RUN go build -o service ./cmd/

FROM alpine:latest

RUN apk add --no-cache iputils curl

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=build /app/service /app/service

EXPOSE 9000

CMD ["/app/service"]