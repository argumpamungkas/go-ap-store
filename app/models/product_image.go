package models

type ProductImage struct {
	GormModel  `gorm:"embedded"`
	ProductID  string `gorm:"size:36;index"`
	Product    Product
	Path       string `gorm:"type:text;"`
	ExtraLarge string `gorm:"type:text;"`
	Large      string `gorm:"type:text;"`
	Medium     string `gorm:"type:text;"`
	Small      string `gorm:"type:text;"`
}
