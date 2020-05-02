FROM golang:latest AS builder
WORKDIR /go
COPY . /go/src/github.com/p-point

RUN go get -u github.com/gin-gonic/gin \
  && go get -u github.com/jinzhu/gorm \
  && go get -u github.com/go-sql-driver/mysql \
  && go get -u bitbucket.org/liamstask/goose/cmd/goose

RUN cd /go/src/github.com/p-point
CMD ["/go/bin/goose","up"]
CMD ["go", "run", "/go/src/github.com/p-point/main.go"]
