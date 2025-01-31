FROM golang:1.23

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod ./

# Download Go modules
RUN go mod download

# Copy the rest of the source code
COPY . .

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["go", "run", "cmd/main.go"]
