FROM golang:1.24.5-alpine AS builder
WORKDIR /app
COPY ./backend .
COPY ../.env .env
RUN go build -o app ./cmd/app/main.go

FROM alpine:3.22 AS production
COPY --from=builder /app/app /app
COPY --from=builder /app/.env .env
EXPOSE 8080:8080
CMD ["./app"]