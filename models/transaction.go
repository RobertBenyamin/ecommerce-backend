package models

import "gorm.io/gorm"

type Transaction struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	ItemID uint
	User   User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Item   Item `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE;"`
}

func MigrateTransaction(db *gorm.DB) {
	db.AutoMigrate(&Transaction{})
}
