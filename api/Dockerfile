# Stage 1: Build
FROM golang:1.23.1 AS builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Create image for running the application
FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
EXPOSE 8080
CMD ["./main"]
