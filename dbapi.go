package main

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
)

func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:vivek@/night_hack?charset=utf8&parseTime=True&loc=Local")
	return db, err
}

func migrate() {

	db, err := connectDB()
	if err != nil {
		fmt.Println("Can not connect to database, Error : ", err)
	}

	fmt.Println("Starting Migration")
	db.AutoMigrate(Product{}, &Category{})
	fmt.Println("Migrated tables Product, Category")                                       //Migrate Product and Category models
	db.Model(&Product{}).AddForeignKey("cat_id", "categories(id)", "RESTRICT", "RESTRICT") //Adding Foreign Key Relationships
	fmt.Println("Added Foreign Key cat_id")
	fmt.Println("Done with migrations")
	defer db.Close()
}

func getCategories() ([]byte, error) {
	db, _ := connectDB()
	var categories []Category
	db.Find(&categories)

	for _, category := range categories {
		fmt.Println(category.ID, category.Name)
	}
	
	defer db.Close()
	result, error := json.Marshal(categories)
	return result, error
}

func createCategory(name string) {
	db, _ := connectDB()
	cat := Category{Name: name}
	db.Create(&cat)
	defer db.Close()
}

func getCategory(categoryId int) ([]byte, error) {
	db, _ := connectDB()
	var category Category
	db.First(&category, categoryId)
	defer db.Close()
	fmt.Println(category)
	result, error := json.Marshal(category)
	return result, error
}

func updateCategory(categoryId int, name string) {
	db, _  := connectDB()
	var category Category
	db.First(&category, categoryId)
	category.Name = name
	db.Save(&category)
	defer db.Close()
}

func deleteCategory(categoryId int) {
	db, _ := connectDB()
	var category Category
	db.First(&category, categoryId)
	db.Delete(&category)
	defer db.Close()
}
