# Start from the official Go image
FROM golang:1.24

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Hot reload
RUN go install github.com/cosmtrek/air@v1.27.3
CMD ["air","-c",".air.toml"]
