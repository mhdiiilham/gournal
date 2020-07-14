FROM golang:1.14

ENV MYSQL_USER="gournal"
ENV APP_DOMAIN="localhost"

RUN mkdir /gournal

ADD . /gournal

WORKDIR /gournal

RUN go build

CMD [ "/gournal/gournal" ]
