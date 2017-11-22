FROM ubuntu:16.04

RUN apt-get update \
 && apt-get -y upgrade \
 && DEBIAN_FRONTEND=noninteractive apt-get -y --no-install-recommends install wkhtmltopdf \
 && apt-get clean

# Install JP font
ADD NotoSerifCJKjp-Light.otf /usr/share/fonts

# MEMO:
# 展開先にフォルダを作って欲しくないので、rootdirのないtargzをわざわざ作成している
# バージョンアップなどの時にはその作業が必須
ADD wkhtmltox-0.12.4_linux-generic-amd64.tar.gz /usr

RUN mkdir /data
WORKDIR /data

