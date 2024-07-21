FROM golang:1.22.5-alpine AS builder

# Install git. Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /pro

COPY ./cmd/main.go /pro
COPY ./config ./pro
COPY ./internal ./pro
COPY ./package ./pro

COPY ./go.mod /pro
COPY ./go.sum /pro

WORKDIR /pro

RUN go get -d -v ./...
# go build -o server main.go
RUN go build -o server

FROM alpine:latest
RUN mkdir /pro
COPY --from=builder /pro/server /pro/server
WORKDIR /pro
CMD ["/pro/server"]


