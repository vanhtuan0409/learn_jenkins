FROM golang:alpine

RUN apk update && \
  apk upgrade && \
  apk add curl

RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

RUN go build -o ./bin/app
RUN ln -s /go/src/app/bin/app /usr/local/bin

CMD ["app"]
