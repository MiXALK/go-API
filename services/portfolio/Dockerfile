ARG GO_VERSION=1.14.2

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /app

COPY . .

RUN apk update \
    && apk add --no-cache git \
    && go mod download \
    && GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE ${API_PORT}

CMD ["./main"]