FROM --platform=linux/x86_64 mysql:5.7

#MySQL設定ファイルをイメージ内にコピー
COPY ./docker/db/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]