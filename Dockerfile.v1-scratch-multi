FROM golang:1.12.1 as builder
WORKDIR /var/tmp
USER daemon
COPY Makefile .
COPY sleepy.go .
ENV GOCACHE=/var/tmp/.gocache
RUN make sleepy
RUN make test

FROM scratch
WORKDIR /
COPY --from=builder /var/tmp/sleepy .
USER 1
ENTRYPOINT ["/sleepy"]
