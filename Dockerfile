FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o academic-tracker-api ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/academic-tracker-api .

EXPOSE 8080

CMD ["./academic-tracker-api"]