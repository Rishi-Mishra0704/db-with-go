package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint
	Name  string
	Price uint
}

func main() {
	dsn := "user=rishi dbname=test_gorm password=1111 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal(err)
	}
	// CRUD Operations

	// Create
	product := Product{Name: "Laptop", Price: 1000}
	product2 := Product{Name: "Mobile", Price: 500}
	result := db.Create(&product)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Created product: %+v\n", product)

	result2 := db.Create(&product2)
	if result2.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Created product: %+v\n", product2)
	// Read
	var fetchedProduct Product
	result = db.First(&fetchedProduct, product.ID)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Found product: %+v\n", fetchedProduct)

	// Update
	db.Model(&fetchedProduct).Update("Price", 1200)
	fmt.Printf("Updated product: %+v\n", fetchedProduct)

	// Delete
	// Delete
	result = db.Delete(&fetchedProduct)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Deleted product: %+v\n", fetchedProduct)

}
