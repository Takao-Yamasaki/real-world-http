FROM golang:1.21-alpine
RUN apk update && apk add curl
ENV ROOT=/go/src/client
WORKDIR ${ROOT}
COPY client/go.mod ${ROOT}/
RUN go mod download

# RUN go mod download \
#     && go build -o main ${ROOT}/client.go
# CMD ${ROOT}/main

# 開発中の場合はこちらを使用する
CMD ["/bin/sh"]
