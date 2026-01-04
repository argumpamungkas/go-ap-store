package fakers

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"

	"github.com/argumpamungkas/go-ap-store-app/app/models"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func ProductFaker(db *gorm.DB) *models.Product {
	// get user dulu
	user := UserFaker(db)
	err := db.Create(&user).Error

	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	s := slug.Make(name)
	slug := strings.ReplaceAll(s, "-", "_")
	sku := slug
	return &models.Product{
		GormModel: models.GormModel{
			ID: uuid.NewString(),
		},
		UserID:           user.ID,
		Sku:              sku,
		Name:             name,
		Slug:             slug,
		Price:            decimal.NewFromFloat(fakePrice()),
		Stock:            rand.Intn(100),
		Weight:           decimal.NewFromFloat(rand.Float64()),
		ShortDescription: faker.Paragraph(),
		Description:      faker.Paragraph(),
		Status:           1,
	}
}

func fakePrice() float64 {
	return precision(rand.Float64()*math.Pow10(rand.Intn(8)), rand.Intn(2)+1)
}

func precision(val float64, pre int) float64 {
	div := math.Pow10(pre)
	return float64(int64(val*div)) / div
}
