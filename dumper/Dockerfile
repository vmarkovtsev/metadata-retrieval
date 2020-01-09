FROM golang:1.13 AS builder
RUN go get -v github.com/athenianco/metadata-retrieval/examples/cmd

FROM postgres:latest

COPY --from=builder /go/bin/cmd /usr/local/bin/metadata-retrieval
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean && rm -rf /var/lib/apt/lists/*
COPY ghsync.sh /usr/local/bin

CMD ["ghsync.sh"]
