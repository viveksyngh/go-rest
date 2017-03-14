package main

import "time"
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"size:255"`
	CatID        uint   `gorm:"size:10"`
	PricePerUnit float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type Category struct {
	ID      uint    `gorm:"size:10"` 
	Name    string  `gorm:"size:255"`
	Product Product `gorm:"ForeignKey:CatID"`
}
