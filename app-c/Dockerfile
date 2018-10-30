FROM alpine as build

RUN apk add --no-cache \
    build-base

COPY . /app/
WORKDIR /app

RUN gcc main.c -o app

CMD ["./app"]
