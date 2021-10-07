FROM golang:1.16-alpine3.13 AS builder

RUN apk update && apk add git

WORKDIR $GOPATH/src/embrio/

COPY . .

ADD https://storage.googleapis.com/bri-embrio-dev_cloudbuild/key/bri-gcs.json /go/src/embrio

ENV GOSUMDB=off
COPY go.mod .
COPY go.sum .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/auth

FROM alpine:3.11

COPY --from=builder /go/bin/auth /go/bin/auth

COPY --from=builder /go/src/embrio/bri-gcs.json /go/src/embrio

RUN apk add --no-cache tzdata

ENTRYPOINT ["/go/bin/auth"]
