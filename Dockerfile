# === BUILDER ===
FROM golang:latest AS builder
WORKDIR /go/src/app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .
RUN go build -o xssfox

# === RUNNER ===
FROM debian:bookworm
RUN mkdir /app

# Copy the binary from the builder stage
COPY --from=builder /go/src/app/xssfox /app/xssfox
COPY --from=builder /go/src/app/samples /app/samples

WORKDIR /app/
CMD ["/app/xssfox"]
