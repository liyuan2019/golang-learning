# Dockerfileの基本構文
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

