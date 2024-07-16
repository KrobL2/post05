# docker build -t goPost .
# docker run -it goPost

FROM golang:1.22.5-alpine AS builder

# Install git. Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /pro
ADD ./main.go /pro/

ADD ./go.mod /pro/
ADD ./go.sum /pro/

WORKDIR /pro
RUN go get -d -v ./...
RUN go build
# go build -o server main.go
FROM alpine:latest

RUN mkdir /pro
COPY --from=builder /pro/server /pro/server
WORKDIR /pro
CMD ["/pro/server"]