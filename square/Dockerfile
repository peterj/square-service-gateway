# Build the Go binary in this first image
FROM golang:1.12.4-alpine AS builder

RUN apk add --no-cache ca-certificates git
WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /app .

# Final image with binary
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs

# Copy the built binary from the previous step
COPY --from=builder /app /app

ENTRYPOINT ["./app"]