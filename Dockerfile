FROM golang:1.21-alpine AS builder

# Install git. Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /restfull-server

WORKDIR /restfull-server

COPY . .
COPY config/main.yaml /restfull-server/config/main.yaml

RUN go mod tidy
# go build -o server main.go
RUN go build -o server ./cmd/main.go

FROM alpine:latest
RUN mkdir /restfull-server
WORKDIR /restfull-server
COPY --from=builder /restfull-server/server ./server
COPY config/main.yaml /restfull-server/config/main.yaml
CMD ["./server"]


