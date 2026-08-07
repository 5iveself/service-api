FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /service-api ./cmd/api

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app/cmd/api
COPY --from=builder /service-api /service-api
COPY --from=builder /app/internal/config/.env /app/internal/config/.env

EXPOSE 8080
CMD ["/service-api"]
