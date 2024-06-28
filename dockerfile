FROM golang:1.22.2
WORKDIR /app
COPY . .
ENV GO111MODULE=on
RUN go build -o dc-bot ./cmd/api/main.go

CMD ["./dc-bot"]