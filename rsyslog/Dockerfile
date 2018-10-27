FROM alpine

RUN apk add --no-cache \
    rsyslog

COPY . /

EXPOSE 514/tcp
EXPOSE 514/udp

ENTRYPOINT ["rsyslogd", "-n", "-f", "/etc/rsyslog.conf"]
