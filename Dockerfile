FROM debian:9

ADD bin/app_linux /go/bin/app

ENTRYPOINT /go/bin/app
WORKDIR /go

EXPOSE 8888
