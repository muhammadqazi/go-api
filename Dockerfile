# Use an official Golang runtime as the base image
FROM golang:1.16 as build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app
RUN go build -o out ./src/cmd

# Final image
FROM gcr.io/distroless/base
COPY --from=build /app/out /

# Set the entrypoint
ENTRYPOINT ["/out"]
