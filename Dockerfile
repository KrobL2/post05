FROM golang:1.21-alpine AS builder

# Install git. Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /go/src/restfull-server

WORKDIR /go/src/restfull-server

COPY . .

RUN go mod tidy
# go build -o server main.go
RUN go build -o server ./cmd/main.go

FROM alpine:latest
RUN mkdir /restfull-server
WORKDIR /restfull-server
COPY --from=builder /go/src/restfull-server/server /go/src/restfull-server/server
CMD ["/go/src/restfull-server/server"]


