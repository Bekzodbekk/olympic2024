FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o country-service ./cmd/main.go

EXPOSE 8003

CMD ["./country-service"]