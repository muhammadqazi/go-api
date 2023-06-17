# Base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o /app/main ./src/cmd/main.go

# Expose the port your application listens on
EXPOSE 9000

# Run the database seed script
RUN chmod +x ./src/scripts/seed-db.sh
RUN ./src/scripts/seed-db.sh

RUN make server
