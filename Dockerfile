FROM golang:1.22.5-alpine AS builder

# Install git. Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /pro
RUN mkdir /pro/config

WORKDIR /pro

COPY config/main.yaml /pro/config

RUN go mod download

ADD ./go.mod /pro/
ADD ./go.sum /pro/

RUN go get -d -v ./...
# go build -o server main.go
RUN go build -o server

FROM alpine:latest
RUN mkdir /pro
COPY --from=builder /pro/server /pro/server
WORKDIR /pro
CMD ["/pro/server"]

