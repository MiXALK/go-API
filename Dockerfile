ARG GO_VERSION=1.14.2

FROM golang:${GO_VERSION}-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

WORKDIR cmd

RUN GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cmd/main .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./main"]
