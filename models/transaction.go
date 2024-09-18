package models

import "gorm.io/gorm"

type Transaction struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	ItemID uint
	User   User `gorm:"foreignKey:UserID"`
	Item   Item `gorm:"foreignKey:ItemID"`
}

func MigrateTransaction(db *gorm.DB) {
	db.AutoMigrate(&Transaction{})
}
