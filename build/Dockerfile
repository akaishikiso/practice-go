FROM golang:1.24.2-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
#  backendディレクトリの作成
RUN mkdir /go/src/backend
# ワーキングディレクトリの設定
WORKDIR /go/src/backend
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/backend