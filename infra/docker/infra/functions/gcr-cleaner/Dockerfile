FROM golang:1.17.2-alpine AS builder

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /go/src/github.com/calmato/shs-web/infra/functions/gcr-cleaner

RUN apk add --update --no-cache \
      git \
      make

ADD ./infra/functions/gcr-cleaner ./
RUN go build -o app ./cmd/server/main.go

FROM alpine

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /var/functions

RUN apk add --update --no-cache \
      ca-certificates

COPY --from=builder /go/src/github.com/calmato/shs-web/infra/functions/gcr-cleaner/app ./app

EXPOSE 8080
CMD ["./app"]
