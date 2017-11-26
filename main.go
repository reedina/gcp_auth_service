package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	a := App{}

	a.InitializeApplication("root", "ls2202!2005", "sre_toolkit", "vscape-187223:us-east1:sre-instance")
	a.RunApplication(":5040")
}
