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
```shell
#コンテナのネットワーク設定
# docker container run [ネットワークオプション] イメージ名[:タグ名] [引数]
# オプション　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　説明
# --add-host=[ホスト名:IPアドレス]　　　　　　　　　　　　　　　　　　　　コンテナの/etc/hostsにホスト名とIPアドレスを定義する
# --dns=[IPアドレス]                                              コンテナ用DNSサーバのIPアドレス指定
# --expose                                                       指定したレンジのポート番号を割り当てる
# --mac-address=[MACアドレス]                                     コンテナのMACアドレスを指定する
# --net=[bridge | none | container:<name | id> | host | NETWORK] コンテナのネットワークを指定する
# --hostname, -h                                                 コンテナ自身のホスト名を指定する
# --publish, -p[ホストのポート番号]:[コンテナのポート番号]              ホストとコンテナのポートマッピング
# --publish-all, -P                                              ホストの任意のポートをコンテナに割り当てる
$ docekr container run -d -p 8080:80 nginx
$ docekr container run -d --dns 192.168.1.1 nginx
$ docker container run -it --add-host test.com:192.168.1.1 centos
```
```shell
#リソースを指定してコンテナを生成/実行
# docker container run [リソースオプション] イメージ名[:タグ名] [引数]
# オプション　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　説明
# --cpu-shares, -C                                             CPUの使用配分(比率)
# --memory, -m                                                 使用するメモリを制限して実行する(単位はb, k, m, gのいずれか)
# --volume=[ホストのディレクトリ]:[コンテナのディレクトリ], -v        ホストとコンテナのディレクトリを共有
$ docker container run --cpu-shares=512 --memory=1g centos
#ディレクトリの共有
$ docker container run -v /Users/li/webap:/usr/share/nginx/html nginx
```
```shell
#コンテナを生成/起動する環境を指定
# docker container run [環境設定オプション] イメージ名[:タグ名] [引数]
# オプション　　　　　　　　　　　　　　　　　　説明
# --env=[環境変数], -e                   環境変数を設定する
# --env-file=[file名]　　　　　　         環境変数をファイルから設定する
# --read-only=[true | false]            コンテナのファイルシステムを読み込む専用にする
# --workdir=[パス], -W                   コンテナの作業ディレクトリを指定する
# --user=[ユーザー名], -u                ユーザー名またUIDを指定する
$ docker container run -it -e foo=bar centos /bin/bash
#環境変数の一括設定
$ docker container run -it --env-file=env_list centos /bin/bash
#作業ディレクトリの設定
$ docker container run -it -w=/tensorflow centos /bin/bash
[root@92d9866d7956 tensorflow]# pwd
/tensorflow
```
```shell
#稼働コンテナ一覧表示
# docker container ls [オプション]
# オプション　　　　　　　　　　　説明
# --all, -a                  起動中/停止中も含めいてすべてのコンテナを表示する
# --filter, -f               表示するコンテナのフィルタリング
# --format                   表示フォーマットを指定
# --last, -n                 最後に起動されてからn件のコンテナのみ表示
# --latest, -l               最後に起動されたコンテナのみ表示
# --no-trunc                 情報を省略しないで表示する
# --quiet, -q                コンテナIDのみ表示
# --size, -s                 ファイルサイズの表示
#起動中/停止中も含めいてすべてのコンテナを表示する
$ docker container ls -a
$ docker container ls -a -f name=test1
$ socker container ls -a -f exited=0
```
```shell
#コンテナの稼働状態確認
$ docker container stats webserver
#コンテナで動作しているプロセス確認
$ docker container top webserver
```
```shell
#コンテナの起動
# オプション　　　　　　　　説明
# --attach, -a         標準出力/標準エラー出力を開く
# --interactive, -i    コンテナの標準入力を開く
$ docker container start [オプション] コンテナ識別子
```
```shell
#コンテナの停止
# --time, -t コンテナの停止時間を指定する(デフォルトは10秒)
#強制的にコンテナを停止する時は、docker container killを使います
$ docker container stop -t 2 dbb4bbe0f470 
```
```shell
#コンテナの再起動
# --time, -t コンテナの再起動時間を指定する(デフォルトは10秒)
$ docker container restart -t 2 webserver
```
```shell
#コンテナの削除
# オプション　　　　　　　　　　　　說明
# --force, -f                 停止中のコンテナを強制的に削除する
# --volumes, -v               割り当てたボリュームを削除する
$ docker container rm webserver
# 停止中の全てコンテナを削除するには
$ docker container prune
```
```shell
#コンテナの中断/再開
$ docker container pause webserver
$ docker container unpause webserver
```
# Dockerコンテナのネットワーク
```shell
# ネットワーク一覧表示
# オプション　　　　　　說明
# -f, --filter=[]    出力をフィルタする
# --no-trunc         詳細を出力する
# -q, --quiet       ネットワークIDのみを表示する
$ docker network ls [オプション]
# 起動したDcokerコンテナの所属するネットワークを確認する
$ docker container inspect コンテナ名
```
```shell
#ネットワークの作成
# オプション　　　　　　　說明
# --driver, -d        ネットワークブリッジまたはオーバレイ(デフォルトはbridge)
# --ip-range          コンテナに割り当てるIPアドレスのレンジを指定
# --subnet            サブネットをCIDR形式で指定
# --ipv6              IPv6ネットワークを有効にするかどうか(true/false)
# -lable              ネットワークに設定するラベル
$ docker network create [オプション] ネットワーク
$ docker network create --driver=bridge web-network
```
```shell
#ネットワークへの接続
# オプション　　　　　　說明
# --ip              IPV4アドレス
# --ip6             IPv6アドレス
# --alias           エイリアス名
# --link            他のコンテナへリンク
$ docker network connect [オプション] ネットワーク コンテナ
$ docker network connect web-network webfront
```
```shell
#コンテナのネットワーク確認
$ docker container inspect コンテナ
```
```shell
# ネットワークを指定したコンテナの起動
$ docker container run -itd --name=webap --net=web-network nginx
```
```shell
#ネットワークからの切断
$ docker network disconnect web-network webfront
```
```shell
#ネットワークの詳細確認
$ docker container inspect [オプション]　ネットワーク
```
```shell
#ネットワークの削除
$ docker network rm [オプション] ネットワーク
```
# 稼働しているDockerコンテナの操作
```shell
# 稼働コンテナへ接続
$ docker container attach コンテナ
# 接続したコンテナごと終了させる時はCtrl+C(効かないみたい)
# コンテナからデタッチする時はCtrl+P Ctrl+Q
```
```shell
#稼働コンテナでプロセス実行
# オプション　　　　　　說明
# --detach, -d      コマンドをバックグラウンドで実行する
# --interactive, -i コンテナの標準入力を開く
# --tty, -t         false    tty(端末デバイス)を使う
# --user, -u        ユーザー名指定
$ docker container exec [オプション]　コンテナ　実行するコマンド [引数]
$ docker container exec -it webserver /bin/echo "Hello world"
```
```shell
#稼働コンテナのプロセス確認
$ docker container top webserver
```
```shell
#稼働コンテナのポート転送確認
$ docker container port webserver
```
```shell
# コンテナの名前変更
$ docker container rename old new
```
```shell
# コンテナ内のファイルをコピー
# docker container cp コンテナ識別子:コンテナ内のファイルパス　ホストのディレクトリパス
$ docker container cp webserver:/etc/nigix/niginx.conf /tmp/niginx
# docker container cp ホストのファイル コンテナ識別子:コンテナ内のファイルパス
$ docker container cp ./test.txt webserver:/tmp/test.txt
```
