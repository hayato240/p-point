FROM golang:latest AS builder

WORKDIR /go
COPY . /go/src/github.com/p-point

RUN go get -u github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql

RUN cd /go/src/github.com/p-point && go build -o bin/p-point src/api/main.go

CMD ["go", "run", "src/api/main.go"]
