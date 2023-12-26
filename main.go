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

type Person struct {
	Name string
	Age  int
}

func main() {
	dsn := "user=rishi dbname=test_gorm password=1111 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Product{}, &Person{})
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

	// Create person
	person := Person{Name: "Rishi", Age: 21}
	result3 := db.Create(&person)
	if result3.Error != nil {
		log.Fatal(result3.Error)
	}
	fmt.Printf("Created person: %+v\n", person)

	// Read Product
	var fetchedProduct Product
	result = db.First(&fetchedProduct, product.ID)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Found product: %+v\n", fetchedProduct)

	// Read Person
	var fetchedPerson Person
	result = db.First(&fetchedPerson, "name = ?", person.Name)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Found person: %+v\n", fetchedPerson)

	// Update
	db.Model(&fetchedProduct).Update("Price", 1200)
	fmt.Printf("Updated product: %+v\n", fetchedProduct)

	// Delete
	result = db.Delete(&fetchedProduct)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Printf("Deleted product: %+v\n", fetchedProduct)
}
