FROM golang:1.20.6-alpine3.18 as builder

ADD . /app

WORKDIR /app

RUN go build .

FROM alpine:3.18.2

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/sample-api .

CMD ["./sample-api"]