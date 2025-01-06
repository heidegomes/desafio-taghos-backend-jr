FROM golang:1.23.4-alpine3.21 as builder

WORKDIR /app

COPY . .

RUN go build -o start main.go

FROM alpine:3.21 as run 

WORKDIR /app

COPY --from=builder /app/start ./start

EXPOSE 8080

CMD ./start
