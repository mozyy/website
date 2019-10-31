FROM mysql:5.6

COPY ./docker/mysql/conf.d /etc/mysql/conf.d

ENV MYSQL_ALLOW_EMPTY_PASSWORD yes