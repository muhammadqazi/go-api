# Use the official Golang image as the base
FROM golang:1.17

# Set the working directory inside the Docker image
WORKDIR /app
RUN pwd && ls -la

# Copy the Go project files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project files
COPY . ./

# Build the Go project using the Makefile
RUN ls
RUN make server

# Expose the necessary port(s) if your application requires it
EXPOSE 9000

# Define the default command to run when the container starts
CMD ["make", "server"]
