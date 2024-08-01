# build stage
FROM golang:latest AS builder

WORKDIR /go/src/coupon-service

COPY . .

WORKDIR /go/src/coupon-service/cmd/coupon_service
RUN go build -o /app/main .

WORKDIR /app
ENTRYPOINT ./main
EXPOSE 8080
