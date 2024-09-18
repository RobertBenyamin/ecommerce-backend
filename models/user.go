package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255;unique"`
	Password string `gorm:"size:255"`
}

func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
