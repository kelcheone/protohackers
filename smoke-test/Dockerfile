# golang TCP server
FROM golang:1.19.3-buster

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# build the go app
RUN go build -o main .

