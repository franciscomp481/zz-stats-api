FROM golang:1.22.5-alpine AS stage1

# Set working directory
WORKDIR /app

# Install CA certificates
RUN apk add --no-cache ca-certificates

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# List contents of /app to verify files are copied
RUN ls -la /app

# Build the Go application, specifying the cmd directory
RUN CGO_ENABLED=0 GOOS=linux go build -o zz-webscrapping-api ./cmd

############################################

FROM scratch

# Copy the built binary and CA certificates from the previous stage
COPY --from=stage1 /app/zz-webscrapping-api /
COPY --from=stage1 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Set the entry point for the container
ENTRYPOINT ["/zz-webscrapping-api"]
