package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	fmt.Println("Go ...")

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
	db.Create(&Product{Code: "L6666", Price: 666})
	db.Create(&Product{Code: "L7777", Price: 777})

	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	fmt.Printf("%v\n", product)

	db.First(&product, "code = ?", "L1212") // find product with code l1212
	fmt.Printf("%v\n", product)

	var product2 Product
	db.First(&product2, "code = ?", "L6666") // find product with code l1212
	fmt.Printf("%v\n", product2)

	db.First(&product, "code = ?", "L7777") // find product with code l1212
	fmt.Printf("%v\n", product)

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}
