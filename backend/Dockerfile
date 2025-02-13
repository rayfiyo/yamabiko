ARG GO_VERSION=1.23
ARG ALPINE_VERSION=3.21

# ==================================================
# ビルドステージ: Goコンパイル環境
# ==================================================
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./cmd/main.go

# ==================================================
# 実行ステージ: 軽量なイメージ
# ==================================================
FROM alpine:${ALPINE_VERSION}

WORKDIR /app
COPY --from=builder /app/server /app/
COPY --from=builder /app/.env /app/.env

# ポート公開
EXPOSE 8080

CMD ["/app/server"]
