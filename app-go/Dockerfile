FROM golang as build

COPY . /go/src/github.com/danielkraic/k8s-rsyslog/app-go
WORKDIR /go/src/github.com/danielkraic/k8s-rsyslog/app-go

RUN CGO_ENABLED=0 GOOS=linux go build -o app .


FROM alpine

COPY --from=build /go/src/github.com/danielkraic/k8s-rsyslog/app-go/app /app/
WORKDIR /app

ENTRYPOINT ./app
