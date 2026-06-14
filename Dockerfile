FROM golang:1.25-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o academic-tracker-api ./cmd/api

EXPOSE 8080

CMD ["./academic-tracker-api"]