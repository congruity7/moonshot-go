FROM golang:1.17.6-alpine3.14  AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=arm go build -o moonshot-go .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./moonshot-go"]