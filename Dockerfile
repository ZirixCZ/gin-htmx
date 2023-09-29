FROM golang:1.20-alpine

WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]