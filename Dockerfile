FROM golang:1.19

WORKDIR /app

COPY . .

VOLUME /app/config

CMD ["go", "run", "main.go", "-f", "config/config.toml"]