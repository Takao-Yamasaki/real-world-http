# 開発環境
FROM golang:1.21-alpine
ENV ROOT=/go/src/server
WORKDIR ${ROOT}
COPY server/go.mod server/go.sum ${ROOT}/
RUN go mod download
CMD ["/bin/sh"]

# 本番環境用
# FROM golang:1.21-alpine
# ENV ROOT=/go/src/server
# WORKDIR ${ROOT}
# COPY server/go.mod server/server.go ${ROOT}/
# RUN go mod download \
#     && go build -o main ${ROOT}/server.go
# RUN chmod 755 ${ROOT}/main
# CMD ${ROOT}/main
