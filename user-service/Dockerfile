FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service ./cmd/main.go

EXPOSE 8001

CMD ["./user-service"]