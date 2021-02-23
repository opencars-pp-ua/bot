FROM golang:1.16-alpine AS build

ENV GO111MODULE=on

WORKDIR /go/src/app

LABEL maintainer="github@shal.dev"

RUN apk add bash git gcc g++ libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /go/bin/bot ./cmd/bot/main.go

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=build /go/bin/bot ./bot
COPY templates/ templates/

ENTRYPOINT ["./bot"]
