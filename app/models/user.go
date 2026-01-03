package models

import (
	"gorm.io/gorm"
)

type User struct {
	GormModel     `gorm:"embedded"`
	Adresses      []Address      // has many
	FirstName     string         `gorm:"size:100;not null"`
	LastName      string         `gorm:"size:100;"`
	Email         string         `gorm:"size:100;not null;uniqueIndex"`
	Password      string         `gorm:"size:255;not null"`
	RememberToken string         `gorm:"size:255;not null"`
	DeletedDate   gorm.DeletedAt // digunakan untuk soft delete / tidak bisa didelete secara permanen
}
