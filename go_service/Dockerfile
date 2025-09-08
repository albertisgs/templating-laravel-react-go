# Base image Golang
FROM golang:latest AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.* ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
# This command ensures the output is a static binary named 'main'
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

# --- Final Stage ---
# Use a minimal alpine image for the final container
FROM alpine:latest

# Copy the built executable from the builder stage
COPY --from=builder /main /main

# Expose port
EXPOSE 8080

# Run the executable
CMD [ "/main" ]