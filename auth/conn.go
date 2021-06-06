package auth

import (
	"fmt"
	"tericsoft/categories"
	"tericsoft/db"
	"tericsoft/products"

	"github.com/jinzhu/gorm"
)

// InitDB creates the connection to databse server and then
// creates the database and the required tables
func InitDB() {
	fmt.Println("intializing database......")
	var err error
	conn := "root:@tcp(localhost:3306)/?parseTime=true"
	db.DB, err = gorm.Open("mysql", conn)

	//uncomment the log mod for sql queries in console
	db.DB.LogMode(true)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.DB.Exec("CREATE DATABASE teric")
	db.DB.Exec("USE teric")

	// Migration to create tables for Users schema
	db.DB.AutoMigrate(&User{}, &products.Products{}, &products.CategoryProduct{}, &categories.Categories{})
}
