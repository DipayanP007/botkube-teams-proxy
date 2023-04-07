FROM golang:1.19.4-alpine as builder

WORKDIR /opt

COPY . .

RUN go mod download

ENV GOOS=linux

RUN go build -o webhook-proxy

# CMD ["/opt/webhook-proxy"]

FROM alpine

WORKDIR /opt/webhook-proxy

COPY --from=builder /opt/webhook-proxy .
COPY --from=builder /opt/.env .

CMD ["/opt/webhook-proxy/webhook-proxy"]