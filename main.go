package main

import (
	"database/sql"
	"log"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var cfg = mysql.Cfg("vscape-187223:us-east1:sre-instance", "root", "ls2202!2005")
	cfg.DBName = "sre_toolkit"
	database, err := mysql.DialCfg(cfg)
	if err != nil {
		log.Fatal("Cannot find database. Received error: " + err.Error())
	} else {
		db = database
		println("works")
	}
	stmt, _ := db.Prepare("INSERT register SET firstname=?,lastname=?,email=?")
	stmt.Exec("mike", "rapuano", "mike@mike")
}
