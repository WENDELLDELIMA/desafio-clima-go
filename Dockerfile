FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o weather-service .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/weather-service .

EXPOSE 8080
ENV PORT=8080
CMD ["./weather-service"]
