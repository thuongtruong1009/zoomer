FROM redis:7.0-alpine

RUN mkdir -p /redis
WORKDIR /redis
RUN apk add --no-cache bash

COPY sentinel.conf ./

EXPOSE 26379

CMD ["redis-server", "/redis/sentinel.conf", "--sentinel"]
