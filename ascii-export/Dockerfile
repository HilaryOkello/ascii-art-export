# Use the official Golang image as a build environment
FROM golang:1.22.6 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod file to the container
COPY go.mod ./

# Download the Go module dependencies
RUN go mod download

# Copy the entire project to the container
COPY . .

# Build the Go application with no CGO (C bindings) enabled
RUN CGO_ENABLED=0 go build -o main .

# Use the lightweight Alpine image for the final stage
FROM alpine:3.18

# Set metadata 
LABEL org.opencontainers.image.title="ascii-art-web" \
    org.opencontainers.image.description="A program that generate ascii-art from a text" \
    org.opencontainers.image.version="1.0" \
    org.opencontainers.image.authors="\
    Quinter Ochieng https://github.com/apondi-art\ 
    John Opiyo  https://github.com/SidneyOps75\
    Hilary Okello https://github.com/HilaryOkello" 

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage to the final image
COPY --from=builder /app/main .

# Copy static files, templates, and ASCII art to the container
COPY static /app/static
COPY templates /app/templates
COPY ascii-art/banner /app/ascii-art/banner

# Expose port 8080 for the application
EXPOSE 8080

# Command to run the application
CMD ["./main"]
