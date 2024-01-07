FROM golang:1.21-alpine
ENV ROOT=/go/src/server
WORKDIR ${ROOT}
COPY server/go.mod server/server.go ${ROOT}/
RUN go mod download \
    && go build -o main ${ROOT}/server.go
RUN chmod 755 ${ROOT}/main
CMD ${ROOT}/main

# 開発中の場合はこちらを使用する
# CMD ["/bin/sh"]
