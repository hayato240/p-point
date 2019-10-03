package infrastructure

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/p-point/interfaces/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

type SqlResult struct {
	Result sql.Result
}

func NewSqlHandler() database.SqlHandler {
	conn, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/p_point")
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{})(database.Result, error){
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res.Result, err
	}
	res.Result = result
	return res.Result, nil
}
