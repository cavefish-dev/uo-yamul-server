# Start from the official Golang image
FROM golang:1.23 AS builder

# Set environment variables
ENV GO111MODULE=on
# Disable CGO to ensure a statically linked binary
ENV CGO_ENABLED=0 
ENV GOOS=linux

# Set the working directory inside the container
RUN mkdir yamul-gateway
WORKDIR /yamul-gateway

# Copy go.mod and go.sum files
COPY ./yamul-gateway/go.mod ./yamul-gateway/go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY ./yamul-gateway .

# Build the Go app
RUN go build -o /main

# Use a minimal image for the final container
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /main .

# Expose port (change if your app uses a different port)
EXPOSE 2593

# Command to run the executable
CMD ["./main"]