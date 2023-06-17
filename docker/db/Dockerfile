FROM --platform=linux/x86_64 mysql:8.0

#MySQL設定ファイルをイメージ内にコピー
COPY ./docker/db/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]