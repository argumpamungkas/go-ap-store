package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	GormModel        `gorm:"embedded"`
	ParentID         string `gorm:"size:36;index"`
	UserID           string `gorm:"size:36;index"`
	User             User
	ProductImages    []ProductImage
	Categories       []Category      `gorm:"many2many:product_categories;"`
	Sku              string          `gorm:"size:100;index"`
	Name             string          `gorm:"size:255;"`
	Slug             string          `gorm:"size:255;"`
	Price            decimal.Decimal `gorm:"typw:decimal(16,2);"`
	Stock            int
	Weight           decimal.Decimal `gorm:"type:decimal(10,2);"`
	ShortDescription string          `gorm:"size:255;"`
	Description      string          `gorm:"type:text;"`
	Status           int
	DeletedDate      gorm.DeletedAt
}
