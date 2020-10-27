FROM golang:1.15 as builder

WORKDIR /app

COPY cmd/ cmd/
COPY go.mod .
COPY go.sum .

RUN CGO_ENABLED=0 go build -o timer cmd/*


FROM alpine

WORKDIR /app

COPY --from=builder /app/timer .

CMD ["./timer"]
