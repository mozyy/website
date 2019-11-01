FROM mysql:5.6

COPY ./docker/mysql/conf.d /etc/mysql/conf.d
COPY ./docker/mysql/docker-entrypoint-initdb.d /docker-entrypoint-initdb.d


ENV MYSQL_ALLOW_EMPTY_PASSWORD yes