package fakers

import (
	"github.com/bxcodec/faker/v3"

	"github.com/argumpamungkas/go-ap-store-app/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// dipanggil oleh seeder untuk di insert ke db
func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		GormModel: models.GormModel{
			ID: uuid.NewString(),
		},
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "password123",
		RememberToken: "",
		DeletedDate:   gorm.DeletedAt{},
	}
}
