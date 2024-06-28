FROM golang:1.22.2

RUN apt-get update && apt-get install -y \
    wget \
    unzip \
    xvfb \
    libxi6 \
    libgconf-2-4 \
    chromium-browser \
    chromium-chromedriver

RUN chmod +x /usr/bin/chromedriver
WORKDIR /app
COPY . .
ENV GO111MODULE=on
RUN go build -o dc-bot ./cmd/api/main.go

CMD ["./dc-bot"]