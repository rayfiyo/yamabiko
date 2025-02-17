# nodeイメージを使ってビルドするステージ
FROM node:18-alpine as build

WORKDIR /app

# package.json / package-lock.json を先にコピー
COPY package*.json ./

# 依存ライブラリのインストール
RUN npm install

# ソースコードをコピーしてビルド
COPY . .
RUN npm run build

# ===============================
# 本番用のNginxイメージに静的ファイルを配置
# ===============================
FROM nginx:stable-alpine

# ビルドで生成された成果物(buildフォルダ)をコピー
COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

