FROM cgr.dev/chainguard/go@sha256:7713389cd43a75360efff7d62f1c75459bdb47c407d6b3bd3ec50bccc6906646 AS builder

WORKDIR /app
COPY . /app

RUN go mod tidy; \
    go build -o main .

FROM cgr.dev/chainguard/glibc-dynamic@sha256:9850190b2e79687e2ffe9948f1648ca780e8a2461dabfb3e275a95f7912f4081

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs docs

ENV ARANGO_HOST localhost
ENV ARANGO_USER root
ENV ARANGO_PASS rootpassword
ENV ARANGO_PORT 8529
ENV MS_PORT 8080

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]
