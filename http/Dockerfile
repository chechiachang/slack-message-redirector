#FROM golang:1.12.1-alpine3.9
FROM golang:1.12.1

ENV SLACK_TOKEN=
ENV SLACK_CHANNEL_ID=

WORKDIR /go/src/github.com/chechiachang/slack-message-redirector/

ADD . .

RUN go get ./... && go build .

CMD ["./slack-message-redirector"]
