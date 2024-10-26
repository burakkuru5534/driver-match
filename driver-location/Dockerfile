# Stage 1 - Build the Go application
FROM golang:1.23 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Install godotenv for loading environment variables
RUN go get github.com/joho/godotenv
# Build the application as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2 - Create a minimal image for running the app
FROM alpine:latest

# Install SSL certificates for any HTTPS requests the app might make
RUN apk --no-cache add ca-certificates

# Set the working directory in the final image
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app will run on
EXPOSE 8081

# Run the application
CMD ["./main"]
