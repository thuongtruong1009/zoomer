FROM redis:7.0-alpine

ENV SENTINEL_DOWN_AFTER 5000
ENV SENTINEL_FAILOVER 5000
ENV SENTINEL_QUORUM 2
ENV SENTINEL_PARALLEL_SYNCS 1

RUN mkdir -p /redis

WORKDIR /redis

RUN apk add --no-cache bash

COPY ./sentinel.conf .
COPY ./sentinel-entrypoint.sh /usr/local/bin/

RUN chown redis:redis /redis/* && chmod +x /usr/local/bin/sentinel-entrypoint.sh

EXPOSE 26379

ENTRYPOINT [ "/usr/local/bin/sentinel-entrypoint.sh" ]
