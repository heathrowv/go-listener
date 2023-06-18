FROM golang:1.20.5-alpine3.18
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-listener


EXPOSE 8080
ENTRYPOINT ["app/go-listener"]