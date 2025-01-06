FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app .

EXPOSE 8080

CMD ["/app/app"]
