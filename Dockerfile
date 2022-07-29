FROM golang:1.17-alpine

ENV GIN_MODE=release

ENV PORT=3004

WORKDIR /go/src/github.com/CalculusEnjoyer/test

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /test .

EXPOSE $PORT

ENTRYPOINT ["/test"]