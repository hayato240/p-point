FROM golang:latest AS builder

WORKDIR /go
COPY . /go/src/github.com/p-point

RUN go get -u github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql

RUN cd /go/src/github.com/p-point
CMD ["go", "run", "/go/src/github.com/p-point/main.go"]
