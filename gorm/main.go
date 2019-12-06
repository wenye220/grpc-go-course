package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	println("Hello world")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to databaes")
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L1212", Price: 1000})

	var product Product
	db.First(&product, 1)

	println(product.Price)
}
