GOPATH:=$(shell go env GOPATH)
export CGO_ENABLED=0
export GOOS=linux

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis:. --micro_out=. --go_out=plugins=grpc:. pb/bookings.proto
	protoc --proto_path=${GOPATH}/src:${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis:. --go_out=plugins=grpc:./gateway --grpc-gateway_out=logtostderr=true:./gateway pb/bookings.proto

.PHONY: build
build: proto

	go build -o bookings-srv main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t bookings-srv:latest
