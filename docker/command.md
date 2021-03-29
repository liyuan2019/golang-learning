```shell
#docker container run Dcokerイメージ名 実行コマンド
$ docker container run ubuntu:latest /bin/echo 'Hello world'
```
```shell
#Dockerのバージョン確認
$ docker version
```
```shell
#Dockerの実行環境確認
$ docker system info
```
```shell
#Dockerのディスク利用状況
$ docker system df
#詳細確認
$ docker system df -v
```
```shell
#イメージのダウンロード
$ docker pull nginx
```
```shell
#イメージの確認
$ docker image ls
```
```shell
#「webserver」という名前のdockerコンテナを起動する
$ docker container run --name webserver -d -p 80:80 nginx
```
```shell
#Dokcerで起動したコンテナの状況を確認
$ docker container ps
```
```shell
#コンテナ稼働確認 docker container stats コンテナ名
$ docker container stats webserver
```
```shell
#コンテナの停止 docker stop コンテナ名
$ docker stop webserver
```
```shell
#コンテナの起動 docker start コンテナ名
$ docker start webserver
```
