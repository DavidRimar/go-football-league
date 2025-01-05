# Use the official Golang image to build the Go app
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY ./consumer .

# Fetch dependencies and build the Go app
RUN go mod tidy
RUN go build -o consumer main.go

# Use a smaller image to run the Go app
FROM alpine:latest

# Install necessary dependencies (like ca-certificates) for RabbitMQ connection
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder image
COPY --from=builder /app/consumer /consumer

# Expose any required ports (optional, since consumer typically doesn’t expose)
EXPOSE 8080

# Set the entry point for the container to run the Go consumer
CMD ["/consumer"]