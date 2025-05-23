FROM golang:1.24.3-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go


FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app /app

ENTRYPOINT ["/app/main"]
