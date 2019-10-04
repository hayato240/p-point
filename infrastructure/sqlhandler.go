package infrastructure

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/p-point/interfaces/database"
	"os"
	"strings"
)

type SqlHandler struct {
	Conn *sql.DB
}

type SqlResult struct {
	Result sql.Result
}

type SqlRow struct {
	Rows *sql.Rows
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}
	return defaultValue
}

func getConnectionString() string {
	host := getParamString("MYSQL_DB_HOST", "mysql")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", "root")
	pass := getParamString("MYSQL_PASSWORD", "root")
	dbname := getParamString("MYSQL_DB", "p_point")
	protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	dbargs := getParamString("MYSQL_DBARGS", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, dbname, dbargs)
}

func NewSqlHandler() database.SqlHandler {
	connectionString := getConnectionString()
	conn, err := sql.Open("mysql", connectionString)

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

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow).Rows, err
	}
	row := new(SqlRow)
	row.Rows = rows

	return row.Rows, nil
}
