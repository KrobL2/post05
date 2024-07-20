FROM golang:1.22.5-alpine AS builder

# Install git. Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /pro

WORKDIR /pro

COPY . .

RUN go mod download

ADD ./cmd/main.go /pro/

ADD ./go.mod /pro/
ADD ./go.sum /pro/

COPY configs/main.yaml /pro/configs/main.yaml

RUN go get -d -v ./...
# go build -o server main.go
RUN go build -o server
FROM alpine:latest

RUN mkdir /pro
COPY --from=builder /pro/server /pro/server
WORKDIR /pro
CMD ["/pro/server"]

