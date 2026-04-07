# Use the official Golang image
FROM golang:1.21-alpine AS build

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Use a minimal image for running
FROM alpine:latest

EXPOSE 8080

WORKDIR /root/

# Copy the binary from the build stage
COPY --from=build /app/main .

# Run the binary
CMD ["./main"]
RUN adduser -u 5678 --disabled-password --gecos "" appuser && chown -R appuser /app
USER appuser

# During debugging, this entry point will be overridden. For more information, please refer to https://aka.ms/vscode-docker-python-debug
CMD ["gunicorn", "--bind", "0.0.0.0:5000", "main:app"]
