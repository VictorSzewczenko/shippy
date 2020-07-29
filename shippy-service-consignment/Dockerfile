FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# Only copy source code now to prevent module layer invalidation.
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-consignment *.go

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/shippy-service-consignment .

CMD ["./shippy-service-consignment"]

# FROM alpine
# ADD shippy-service-consignment /shippy-service-consignment
# ENTRYPOINT [ "/shippy-service-consignment" ]