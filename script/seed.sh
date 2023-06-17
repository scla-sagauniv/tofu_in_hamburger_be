#!/bin/bash

# シードデータのSQLファイルパス
current_dir=$(pwd)
SEED_SQL_FILE="$current_dir/seed_data.sql"

# MySQLに接続してシードデータを実行
mysql -h "$MYSQL_HOST" -P "$MYSQL_PORT" -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" "$MYSQL_DATABASE" < "$SEED_SQL_FILE"

echo "Seeding completed."
