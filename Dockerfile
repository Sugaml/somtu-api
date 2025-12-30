# Stage 1: Build the Go binary with CGO enabled
FROM golangci/golangci-lint:v1.63 AS base
LABEL maintainer="tamangsugam09@gmail.com"
LABEL last_modified="2024-09-01"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy

# Stage 2: Build
FROM base AS builder
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Stage 5: Final image with just the binary
FROM alpine AS final
WORKDIR /app
RUN touch .env
COPY --from=builder /app/main .
COPY --from=builder /app/docs/* ./docs/
CMD ["/app/main"]
