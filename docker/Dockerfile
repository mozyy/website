# FROM alpine/git:1.0.7 AS repo

# RUN git clone --depth=1 -b gh-pages https://github.com/mozyy/website.git
# RUN ls /git/website


FROM nginx:1.16.0
COPY ./nginx/conf.d /etc/nginx/conf.d
COPY ./nginx/ssl /etc/nginx/ssl
COPY ./nginx/www /usr/share/nginx/www
# COPY --from=repo /git /usr/share/nginx/www
VOLUME /var/log/nginx
EXPOSE 443