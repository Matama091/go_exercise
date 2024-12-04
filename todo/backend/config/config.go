package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
}
