package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/p-point/interfaces/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() database.SqlHandler {
	conn, err := sql.Open("mysql", "root:root@tcp(mysql-container:3306)/p_point")
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}