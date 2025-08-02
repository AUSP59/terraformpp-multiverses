FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod init omniforge && go mod tidy

RUN go build -o omniforge ./cmd/omniforge/main.go

CMD ["./omniforge", "status"]