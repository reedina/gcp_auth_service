package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	a := App{}

	//a.InitializeApplication("root", "ls2202!2005", "sre_toolkit", "vscape-187223:us-east1:sre-instance")
	a.InitializeApplication(os.Getenv("AUTH_API_DB_USER"),
		os.Getenv("AUTH_API_DB_PASSWORD"),
		"sre_toolkit",
		os.Getenv("AUTH_API_DB_HOST"))

	a.RunApplication(":5040")
}

/*
WORDPRESS_DB_HOST

WORDPRESS_DB_USER

WORDPRESS_DB_PASSWORD
*/
