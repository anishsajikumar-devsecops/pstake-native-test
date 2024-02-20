# Use Go 1.19 on Alpine for smaller image size and build environment
FROM golang:1.19-alpine AS build-env

# Install build dependencies
RUN apk add --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/persistenceOne/pstake-native

# Copy the go.mod and go.sum files first to cache dependencies
COPY go.mod go.sum ./
# Download Go modules before copying the entire source code to leverage Docker cache
RUN go mod download

# Add the rest of the source files from the repository
COPY . .

# Build the project - consider using specific make targets if available
RUN make install

# Use a small base image for the runtime
FROM alpine:edge

# Install runtime dependencies
RUN apk add --no-cache ca-certificates jq bash curl perl

# Set working directory
WORKDIR /root

# Copy the built binary from the build-env
COPY --from=build-env /go/bin/pstaked /usr/bin/pstaked

# Document that the service listens on these ports.
EXPOSE 1317 26656 26657 9090

# Set the default command to run the binary
CMD ["/usr/bin/pstaked", "version", "--long"]
