# Stage 1: Build
FROM golang:1.23.1 AS builder
WORKDIR /app
COPY . ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o consumer .

# Stage 2: Create image for running the application
FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/consumer .
EXPOSE 8080
CMD ["./consumer"]