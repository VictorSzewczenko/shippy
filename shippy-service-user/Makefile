
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=. --go_out=. --micro_out=. proto/user/user.proto

.PHONY: build
build: proto
	go build -o shippy-service-user *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t shippy-service-user:latest

.PHONY: docker-push
docker-push:
	docker tag shippy-service-user:latest victorszewczenko/shippy-service-user:latest
	docker push victorszewczenko/shippy-service-user:latest