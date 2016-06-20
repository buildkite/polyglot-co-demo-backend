FROM golang:1.6

EXPOSE 8080

ADD . /go/src/github.com/buildkite/polyglot-co-demo-backend

WORKDIR /go/src/github.com/buildkite/polyglot-co-demo-backend

CMD ["go","run","main.go"]
