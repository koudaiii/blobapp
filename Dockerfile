# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add environment variables
ENV AZURE_STORAGE_ACCOUNT_NAME=value

# Set the Current Working Directory inside the container
WORKDIR /app

# Add files to app folder
ADD . /app

# Build the Go app
RUN go build -o bin/blobapp .

# Expose ports to the outside world
EXPOSE 1323

# Command to run the executable
CMD ["/app/bin/blobapp"]
