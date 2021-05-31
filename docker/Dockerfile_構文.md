## Dockerfileの基本構文
```
命令　引数
```
Dockerfileの命令
|命令|說明|
|:----|:----|
|FROM|ベースイメージの指定|
|RUN|コマンド実行|
|CMD|コンテナの実行コマンド|
|LABEL|ラベルを設定|
|EXPOSE|ポートのエクスポート|
|ENV|環境変数|
|ADD|ファイル/ディレクトリの追加|
|COPY|ファイルのコピー|
|ENTRYPOINT|コンテナの実行コマンド|
|VOLUME|ボリュームのマウント|
|USER|ユーザーの指定|
|WORKDIR|作業ディレクトリ|
|ARG|Dockerfile内変数|
|ONBUILD|ビルド完了後に実行される命令|
|STOPSIGNAL|システムコールシグナルの設定|
|HEALTHCHECK|コンテナのヘルスチェック|
|SHELL|デフォルトシェルの設定|

## ベースイメージ(FROM命令)
```
FROM [イメージ名]
FROM [イメージ名]:[タグ名]
FROM [イメージ名]@[ダイジェスト]
```
## DockerfileからDockerイメージの作成(docker build)
```shell
$ docker build -t [生成するイメージ名]:[タグ名] [Dockerfileの場所]
```
標準入力からのビルド
Dockerfileだけでビルドに必要なファイルを含めることができないため、Dockerfileと必要なファイルをtarでまとめて標準入力から指定します。
```shell
$ docker build - < Dockerfile
$ docker build - < docker.tar.gz
```
## マルチステージビルド
DockerにはアプリケーションをビルドするためのDockerイメージと、プロダクションで実際に動作させるDockerイメージを同時に作成できる機能があります。
```
FROM golang AS bulider
```
## RUN命令
#### shell形式
```
RUN apt-get insall -y niginx
```
#### Exec形式
シェルを介さず直接実行します。
```shell
RUN ["bin/bash","-c","apt-get install -y niginx"]
```
```
# ベースイメージ設定
FROM ubuntu:latest

# RUN命令実行
RUN echo "shell形式です"
RUN ["echo","Exec形式です"]
RUN ["bin/bash","-c","echo 'Exec形式でbashを使う'"]
```
Dockerfileをビルドすると、記述された１命令ごとに、内部イメージが１つ作成されます。そのため、Dockerfileの命令を減らす工夫がいくつあります。
```
RUN yum -y install\
           httpd\
           php\
           php-mbstring\
           php-pear

```