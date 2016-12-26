FROM golang

ADD . /go/src/github.com/aliaksei-kasiyanik/places-api
ADD ./config.prod.json /go/src/github.com/aliaksei-kasiyanik/places-api/config.json

WORKDIR /go/src/github.com/aliaksei-kasiyanik/places-api

RUN go get

RUN go install github.com/aliaksei-kasiyanik/places-api

ENTRYPOINT /go/bin/places-api

EXPOSE 8080