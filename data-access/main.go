package main

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql/"
)

var db *sql.DB

func main() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
}
