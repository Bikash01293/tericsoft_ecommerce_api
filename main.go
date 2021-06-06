package main

import (
	"fmt"

	"tericsoft/auth"
	"tericsoft/routes"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func main() {
	fmt.Println("Starting server Connection.....")
	//Initializes the database
	auth.InitDB()
	// Routes shows all the end points of the router
	routes.Routes()
}
