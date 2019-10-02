package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/gin-gonic/gin"
)

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("P_POINT_DB_USER")
	PASS := os.Getenv("P_POINT_DB_PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("P_POITN_DOMAIN") + ":" + os.Getenv("P_POINT_PORT") + ")"
	DBNAME := os.Getenv("DBNAME") + "?parseTime=true&loc=Asia%2FTokyo"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connected: ", &db)
	return db
}

func main() {
//	db, err := gorm.Open("mysql", "root@/testdb?charset=utf8&parseTime=True")
	//fmt.Println(err)

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	r.Run(":8080")
//	defer db.Close()
//	db.LogMode(true)
}

