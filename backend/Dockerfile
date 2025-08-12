# build
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o backend ./cmd/moxer

# run
FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/backend .
EXPOSE 8080
CMD ["./backend"]