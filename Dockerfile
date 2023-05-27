# Use a lightweight Go language base image
FROM golang:1.17-alpine

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Install CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon

# Expose the desired port
EXPOSE 8080

# Set the command to run the application with CompileDaemon and environment variables
CMD ["CompileDaemon", "-color=true", "-command=./main", "-log-prefix=false", "-graceful-kill=true"]