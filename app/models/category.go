package models

type Category struct {
	GormModel `gorm:"embedded"`
	ParentID  string `gorm:"size:36;"`
	SectionID string `gorm:"size:36;index"`
	Section   Section
	ProductID string    `gorm:"size:36;index"`
	Product   []Product `gorm:"many2many:product_categories;"` // bisa mempunyai banyak product
	Name      string    `gorm:"size:100;"`
	Slug      string    `gorm:"size:100;"`
}
