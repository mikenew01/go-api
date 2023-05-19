# Build stage
FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o api .

# Final stage
FROM alpine:latest

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release

RUN apk --no-cache add ca-certificates

WORKDIR /opt/

# Copy the pre-built binary from the previous stage
COPY --from=builder /build/api .

# Expose port
EXPOSE 8080

# Command to run
CMD ["./api"]
