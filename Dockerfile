FROM golang:1.21.6 AS builder

WORKDIR /go/bin

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .


RUN go build -o /go/bin/dictionary  *.go


FROM ubuntu

COPY --from=builder /go/bin/dictionary /go/bin/dictionary
COPY --from=builder /go/bin/config/ /go/bin/config/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


WORKDIR /go/bin/

EXPOSE 5300

ENTRYPOINT ["/go/bin/dictionary"]
