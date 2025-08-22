FROM agan-harbor-registry.cn-hangzhou.cr.aliyuncs.com/saas-public/golang-alpine:1.25.0 AS builder
ENV GO111MODULE=on \
    APP_ENV=prod \
    GOPROXY=https://goproxy.cn\
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    TZ="Asia/Shanghai" \
    GIN_MODE=release

WORKDIR /workdir
COPY . .

RUN go mod tidy
RUN go build -o main .

FROM agan-harbor-registry.cn-hangzhou.cr.aliyuncs.com/saas-public/alpine:edge

WORKDIR /workdir
COPY --from=builder /workdir/main .
COPY --from=builder /workdir/config ./config
CMD ["/workdir/main"]

