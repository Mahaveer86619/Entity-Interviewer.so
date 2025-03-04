FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go-admin ./cmd/main.go

FROM gcr.io/distroless/base-debian10
COPY --from=builder /go-admin /go-admin
COPY static /static

EXPOSE 8080

CMD ["/go-admin"]