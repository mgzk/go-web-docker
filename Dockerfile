FROM golang:1.14

RUN go get -u github.com/google/uuid

RUN mkdir /go-web-docker
ADD . /go-web-docker
WORKDIR /go-web-docker

RUN go build -o main .

CMD ["/go-web-docker/main"]