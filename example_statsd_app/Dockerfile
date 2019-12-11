FROM golang:alpine AS builder

LABEL maintainer="Ikhsan Faruqi <faruqisan@gmail.com>"

WORKDIR /app
ADD . /app
RUN cd /app & go mod download
RUN cd /app & go build -o example main.go

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /app/example /app

ENTRYPOINT ./example