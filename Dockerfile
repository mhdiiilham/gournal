FROM golang:1.14-alpine
RUN mkdir /gournal
ADD . /gournal
WORKDIR /gournal
RUN go build
CMD [ "/gournal/gournal" ]
