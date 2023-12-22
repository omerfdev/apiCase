# Use the official Golang image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./main"]

