FROM node:12.7.0 AS builder

RUN git clone --depth=1 -b master https://github.com/mozyy/website.git && \
    cd /website/white/ && \
    yarn && \
    yarn build

FROM nginx:1.16.0
# COPY ./conf/nginx.conf /etc/nainx/nginx.conf
COPY --from=builder /website/white/build /usr/share/nginx/html
# VOLUME ["/var/log/nginx"]
