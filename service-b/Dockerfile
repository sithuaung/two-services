FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]

