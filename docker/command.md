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
*********************************************************
# Dockerイメージの操作

```shell
#イメージのダウンロード
$ docker image pull [オプション] イメージ名[:タグ名]
```
```shell
#イメージ一覧表示
# オプション　　　説明
# -all, -a     すべてのイメージを表示
# --digests    ダイジェストを表示するかどうか(摘要)
# --no-trunc   結果をすべて表示する  
# --quiet, -q  DockerイメージIDのみ表示
$ docker image ls [オプション] [レポジトリ名]
```
```shell
#DCT機能の有効化(ダウンロードする時イメージの検証)
$ export DOCKER_CONTENT_TRUST=1
```
```shell
#イメージの詳細確認
$ docker image inspect [イメージ名]
#--formatオプションで、JSON形式データ指定する
$ docker image inspect --format="{{ .Os}}" nginx
```
```shell
#イメージのタグ設定(イメージに別名をつける)
# docker image tag [元イメージ名] [<Docker Hubのユーザー名>/イメージ名:タグ名]
$ docker image tag nginx liyuan2021/webserver:1.0
```
```shell
#イメージの検索
# docker search [オプション] 検索キーワード
# オプション　　　       説明
# --no-trunc          結果をすべて表示する
# --limit             n件の検索結果を表示する
# --filter=stars=n    お気に入りの数(n以上)の指定
$ docker search --filter=stars=1000 nginx
```
```shell
#イメージの削除
# docker image rm [オプション] [イメージ名/IMAGE ID]
# [IMAGE ID]は先頭3桁程度でいいです
#複数のイメージ名をスペーズで区切り
# オプション　　　　説明
# --force, -f    イメージを強制的に削除する
# --no-prun      中間イメージを削除しない
$ docker image rm -f b8c
```
```shell
#未使用のDockerイメージを削除する
#定期的に削除すると良い
# docker image prune [オプション]
# オプション　　　　　説明
# --all, a         使用していないイメージをすべて削除
# --force, -f      イメージを強制的に削除する
$ docker image prune -a
```
```shell
#Docker Hubへのログイン
# docker login [オプション] [サーバ]
# オプション            説明
# --password, -p      パスワード
# --usename, -u       ユーザー名
$ docker login   
```
```shell
#イメージのアップロード
$ docker image push <Docker Hubユーザー名>/イメージ名:[タグ名]
```
```shell
#Docker Hubからログアウト
$ docker logout [サーバ名]
```
*********************************************************
# Dockerコンテナの操作
```
Dockerコンテナのライフサイクル
コンテナ生成(docker container create)
↓↓↓
コンテナ生成/起動(docker container run)
↓↓↓
停止中コンテナ起動(docker container start)
↓↓↓
起動しているコンテナ停止(docker container stop)
↓↓↓
コンテナ削除(docker container rm)
```

```shell
# コンテナ生成/起動
# docker container run [オプション]　イメージ名[:タグ名] [引数]
# オプション　　　　　　　説明
# --attach, -a        標準入力(STDIN)/標準出力(STDOUT)/標準エラー出力(STDERR)にアタッチする
# --cidfile           コンテナIDをファイルに出力する
# --detach, -d        コンテナを生成し、バックグラウンドで実行する
# --interactive, -I   コンテナの標準入力を開く
# --tty, -t           端末デバイスを使う
$ docker container run -it --name "test1" centos /bin/cal
# -it コンソールに結果を出すオプション(-i -t)
# --name コンテナ名
$ docker container run -it --name "test2" centos /bin/bash
# コンテナ内でシェルを実行する
#exitでシェルを終了する
```
```shell
#コンテナのバックグラウンド実行
# docker container run [実行オプション]　イメージ名[:タグ名] [引数]
# オプション　　　　　　　　　                                                     説明
# --detach, -d                                                                バックグラウンドで実行する
# --user, -u                                                                  ユーザー名を指定
# --restart=[no | on-failure | on-failure:回数n | always | unless-stopped]     コマンドの実行結果によって再起動を行うオプション
# --rm                                                                        コマンドを実行完了後にコンテナを自動で削除
$ docker container run -d centos /bin/ping localhost
```
```shell
#コンテナのログ確認
#-tオプションでタイムスタンプを表示する
$ docker container logs -t 4b47697516a7
```