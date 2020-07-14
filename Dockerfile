FROM golang:1.14

ENV MYSQL_USER="gournal"
ENV MYSQL_PORT="3306"
ENV MYSQL_HOST="database-1.c8gtlmqhel5b.us-east-1.rds.amazonaws.com"
ENV MYSQL_USER_PASSWORD="root"
ENV APP_DOMAIN="localhost"

RUN mkdir /gournal

ADD . /gournal

WORKDIR /gournal

RUN go build

CMD [ "/gournal/gournal" ]
