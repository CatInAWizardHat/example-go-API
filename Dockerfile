# Use latest go image
FROM golang:1.25.0-alpine3.21 AS builder

# Create working directory
WORKDIR /server

# Copy go.mod, go.sum  to working directory
COPY go.mod go.sum ./

RUN go mod download

# Copy src to working directory
COPY . .

# Build
RUN GOOS=linux GOARCH=amd64 go build -o server ./cmd/api

FROM scratch

COPY --from=builder /server/server .

# Pick port
EXPOSE 8080

# Start
CMD ["/server"]
