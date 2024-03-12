# syntax=docker/dockerfile:1

FROM golang:1.22.1-alpine3.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /build

# Run
CMD ["/build"]