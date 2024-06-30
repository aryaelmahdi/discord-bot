FROM golang:1.22.2

# Set the working directory
WORKDIR /app

# Copy the application files
COPY . .

# Set environment variables
ENV GO111MODULE=on

# Install dependencies and Chrome + ChromeDriver
RUN apt-get update && \
    apt-get install -y wget unzip gnupg2 && \
    wget -q -O google-chrome-stable_current_amd64.deb https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && \
    apt-get install -y ./google-chrome-stable_current_amd64.deb || apt-get install -y -f && \
    rm google-chrome-stable_current_amd64.deb && \
    wget -O /tmp/chromedriver.zip https://storage.googleapis.com/chrome-for-testing-public/126.0.6478.126/linux64/chromedriver-linux64.zip && \
    unzip /tmp/chromedriver.zip -d /usr/bin/ && \
    chmod +x /usr/bin/chromedriver-linux64/chromedriver && \
    ln -s /usr/bin/chromedriver-linux64/chromedriver /usr/bin/chromedriver && \
    rm /tmp/chromedriver.zip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Build the Go application
RUN go build -o discord-bot ./cmd/api/main.go

# Expose the port for ChromeDriver
EXPOSE 4444

# Start ChromeDriver and the Go application
CMD chromedriver --allowed-ips --port=4444 & ./discord-bot