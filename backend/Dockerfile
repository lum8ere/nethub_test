# Stage 1: Build phase
FROM golang:1.24-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /nethub-api ./cmd/api/

# Stage 2: Final image
FROM alpine:latest

WORKDIR /app

# RUN apk add --no-cache ca-certificates tzdata
# Копируем только бинарник
COPY --from=builder /nethub-api .
# Копируем папку docs для работы Swagger
COPY --from=builder /app/docs ./docs
EXPOSE 9000
CMD ["./nethub-api"]