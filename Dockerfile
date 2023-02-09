FROM golang:1.19

WORKDIR /cli

COPY . .

RUN go get -d -v ./...

RUN go build -o main .

CMD ["./main -f ./config/config.toml"]