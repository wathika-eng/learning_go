FROM docker.io/library/golang:1.23.0 AS builder

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

# podman run -it -p 8080:8080 localhost/go-app:latest
