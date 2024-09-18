package models

import "gorm.io/gorm"

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	Price       float64
	Description string `gorm:"size:500"`
	UserID      uint
}

func MigrateItem(db *gorm.DB) {
	db.AutoMigrate(&Item{})
}
