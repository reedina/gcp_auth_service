package model

import (
	"database/sql"
	"log"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"

	//Initialize mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//ConnectDB - Establish DB connection
func ConnectDB(user, password, dbname, url string) {

	cfg := mysql.Cfg(url, user, password)
	cfg.DBName = dbname
	database, err := mysql.DialCfg(cfg)
	if err != nil {
		log.Fatal("Cannot find database. Received error: " + err.Error())
	} else {
		db = database
		println("Connected to Database")
	}

}
