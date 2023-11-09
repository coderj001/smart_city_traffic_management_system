FROM golang:alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Create app directory
RUN mkdir /app

# set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Install Air using the recommended install script
RUN go install github.com/cosmtrek/air@latest

# Verify Air installation
RUN air -v

# Copy the rest of your application code
COPY . .

# Expose port 8000
EXPOSE 8000

# Set the entrypoint to `air`. Note the full path to the air binary.
ENTRYPOINT ["air", "-c", "air.toml"]

# Health check
HEALTHCHECK --interval=30s --timeout=3s \
  CMD wget --quiet --tries=1 --spider http://localhost:8000/health || exit 1
