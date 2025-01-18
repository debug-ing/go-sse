FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd
RUN go build -o /app/main main.go

FROM alpine:latest

LABEL maintainer="Mohammad Mahdi Mohammadi <m.m.mohamadi1997@gmail.com>"

WORKDIR /app/

COPY --from=builder /app/main .

EXPOSE 8080

ENTRYPOINT ["./main"]