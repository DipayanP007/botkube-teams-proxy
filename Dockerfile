FROM golang:1.19.4-alpine as builder

WORKDIR /opt

COPY . .

RUN go mod download

ENV GOOS=linux

RUN go build -o botkube-teams-proxy


FROM alpine

WORKDIR /opt/webhook-proxy

COPY --from=builder /opt/botkube-teams-proxy .
COPY --from=builder /opt/.env .

CMD ["/opt/webhook-proxy/botkube-teams-proxy"]