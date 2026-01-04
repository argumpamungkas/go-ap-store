package seeders

import (
	"github.com/argumpamungkas/go-ap-store-app/databases/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

// register seeder
func RegisterSeeder(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.UserFaker(db)},
		{Seeder: fakers.ProductFaker(db)},
	}
}

// proses create seeder ke db
func DBSeed(db *gorm.DB) error {
	for _, value := range RegisterSeeder(db) {
		err := db.Debug().Create(value.Seeder).Error
		if err != nil {
			return err
		}
	}

	return nil
}
