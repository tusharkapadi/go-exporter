# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /Users/tushar.kapadi/Code/go-exporter
COPY go.mod ./
COPY go.sum ./
RUN go mod download



COPY *.go ./


RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/promauto
RUN go get github.com/prometheus/client_golang/prometheus/promhttp
RUN go build -o /exporter

EXPOSE 2112

CMD ["/exporter"]