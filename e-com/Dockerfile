# Use an official Golang runtime as a parent image
FROM golang:1.17-alpine AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the source code and config file into the container at /app
COPY . .

# Build the Go app
RUN go build -o main .

# Use an official Alpine Linux image as a parent image
FROM alpine:latest

# Copy the built executable and config file from the previous stage
COPY --from=builder /app/main /app/main
#COPY --from=builder /app/config.yaml /app/config.yaml

# Set the working directory to /app
WORKDIR /app

# Expose port 8080
EXPOSE 8080

# Run the command to start the app with the config file
CMD ["/app/main"]