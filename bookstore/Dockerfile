# # Build stage
# FROM golang:1.18.4-alpine3.16 AS builder
# WORKDIR /app
# COPY . .
# RUN go build -o main main.go

# # Run stage
# FROM alpine:3.16
# WORKDIR /app
# COPY --from=builder /app/main .

# EXPOSE 8080
# CMD ["/app/main"]



# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the source code and config file into the container at /app
COPY . .
RUN go mod download
# Build the Go app
RUN go build -o main .

# Use an official Alpine Linux image as a parent image

# Expose port 8080
EXPOSE 8080

# Run the command to start the app with the config file
CMD ["/app/main"]