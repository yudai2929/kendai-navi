#!/bin/bash

# 環境変数を設定する
# データベースのユーザー名
DB_USER=${DB_USER:-postgres}
# データベースのパスワード
DB_PASS=${DB_PASS:-password}
# データベースのホスト名
DB_HOST=${DB_HOST:-localhost}
# データベースのポート
DB_PORT=${DB_PORT:-5432}
# データベースの名前
DB_NAME=${DB_NAME:-app}
# SSL 接続を有効にするかどうか
DB_SSL_MODE=${DB_SSL_MODE:-disable}

# マイグレーションファイルが存在するディレクトリのパス
MIGRATIONS_DIR=${MIGRATIONS_DIR:-"file://db/migrate/migrations"}

# Atlas コマンドを実行
atlas migrate apply \
  --dir "$MIGRATIONS_DIR" \
  --url postgres://$DB_USER:$DB_PASS@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=$DB_SSL_MODE