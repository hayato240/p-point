FROM golang:latest AS builder

WORKDIR /go
COPY . /go/src/github.com/hayato240/p-point

RUN go get -u github.com/gin-gonic/gin \
  && go get -u github.com/jinzhu/gorm \
  && go get -u github.com/go-sql-driver/mysql \
  && go get -u github.com/rubenv/sql-migrate \
  && go get -u github.com/pkg/errors

RUN cd /go/src/github.com/hayato240/p-point
CMD ["go", "run", "/go/src/github.com/hayato240/p-point/main.go"]
