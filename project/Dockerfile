FROM golang:1.23.0

WORKDIR /app

COPY go.* ./    

RUN go mod download

COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]