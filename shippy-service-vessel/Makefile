
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=. --go_out=. --micro_out=. proto/vessel/vessel.proto

.PHONY: build
build: proto
	go build -o shippy-service-vessel *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t shippy-service-vessel:latest

.PHONY: docker-push
docker-push:
	docker tag shippy-service-vessel:latest victorszewczenko/shippy-service-vessel:latest
	docker push victorszewczenko/shippy-service-vessel:latest