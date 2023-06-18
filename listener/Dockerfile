FROM golang:1.20.5-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-listener



FROM alpine:3.18.2
WORKDIR /app/
COPY --from=builder /app/go-listener .
EXPOSE 8080
ENTRYPOINT ["app/go-listener"]