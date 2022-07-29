FROM golang:1.17-alpine

ENV GIN_MODE=release

ENV PORT=3004

WORKDIR /go/src/github.com/CalculusEnjoyer/btc-rate-email-sender

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /btc-rate-email-sender .

EXPOSE $PORT

ENTRYPOINT ["/btc-rate-email-sender"]