package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "os"
    "strings"
)

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

func Connection() *gorm.DB {
    connectionString := getConnectionString()
    db, err := gorm.Open("mysql", connectionString)
    if err != nil {
        panic("failed to connect database")
    }
    db.LogMode(true)
    return db
}