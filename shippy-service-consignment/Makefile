
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=. --go_out=. --micro_out=. proto/consignment/consignment.proto

.PHONY: build
build: proto
	go build -o shippy-service-consignment *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t shippy-service-consignment:latest