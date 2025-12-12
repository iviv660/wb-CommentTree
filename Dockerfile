# ---------- builder ----------
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o comment-tree ./cmd/app

# ---------- runner ----------
FROM alpine:3.20

WORKDIR /app

RUN adduser -D appuser
USER appuser

COPY --from=builder /app/comment-tree /app/comment-tree
COPY .env /app/.env

EXPOSE 8080

CMD ["/app/comment-tree"]
