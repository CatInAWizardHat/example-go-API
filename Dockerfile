# Use latest go image
FROM golang:latest

# Create working directory
WORKDIR /server

# Copy go.mod, go.sum  to working directory
COPY go.mod go.sum ./

# Copy src to working directory
COPY . .

# Build
RUN go build -o server

# Pick port
EXPOSE 8080

# Start
CMD ["/server/server"]
