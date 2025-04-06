FROM golang:1.24.2-alpine
# アップデートとgitのインストール
RUN apk update && apk add git
# practice-goディレクトリの作成
RUN mkdir /go/src/practice-go
# ワーキングディレクトリの設定
WORKDIR /go/src/practice-go
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/practice-go