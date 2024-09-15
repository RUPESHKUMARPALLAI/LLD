FROM  golang:1.20.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y libc6

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
