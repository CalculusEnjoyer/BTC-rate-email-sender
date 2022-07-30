FROM golang:1.17-alpine

ENV GIN_MODE=release

ENV PORT=80

WORKDIR /myapp/btc-rate

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /btc-rate .

EXPOSE $PORT

ENTRYPOINT ["/btc-rate"]