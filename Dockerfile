FROM golang:1.21-alpine as base
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY go.mod server.go ${ROOT}/
RUN go mod download \
    && go build -o server ${ROOT}/server.go
CMD ${ROOT}/server

# 開発中の場合はこちらを使用する
# CMD ["/bin/sh"]
