FROM alpine:3.20

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

WORKDIR /root/

ADD https://github.com/pressly/goose/releases/download/v3.22.1/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

ADD migrations/*.sql migrations/
ADD migration_prod.sh .
ADD prod.env .
RUN chmod +x migration_prod.sh

ENTRYPOINT ["bash","migration_prod.sh"]


