# Alpine Linuxとは、軽量なLinuxディストリビューション
# golang:1.20-alpineとは、golangの1.20をAlpine Linuxで動かすためのイメージ
# 通常のgolangイメージは、Debianベースのため、サイズが大きい
FROM golang:1.20-alpine

# apkとは、Alpine Linuxのパッケージ管理ツール
# Alpine Linuxには最初は何も入っていないので、まずはupdateを行う
# その後、bashとgitをインストール
# 最後に、キャッシュを削除
RUN apk update && \
    apk add bash git && \
    rm -rf /var/cache/apk/*

# 作業ディレクトリを/appに指定
WORKDIR /app

# ホストのgo.modとgo.sumをコンテナの作業ディレクトリにコピー
COPY go.mod go.sum ./
RUN go mod download

# ホストのファイルをコンテナの作業ディレクトリにコピー
COPY . .

# 環境変数を設定
# CGO_ENABLEDを0にすることで、cgoを無効化
# GOOSをlinuxにすることで、Linux用のバイナリをビルド
ENV CGO_ENABLED=0 \
    GOOS=linux 

# ビルド
RUN go build -o /portfolio-server command/main.go

# ビルドしたバイナリを実行
CMD ["/portfolio-server"]
