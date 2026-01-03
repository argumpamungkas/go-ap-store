package models

type Section struct {
	GormModel  `gorm:"embedded"`
	Name       string `gorm:"size:100;"`
	Slug       string `gorm:"size:100;"`
	Categories []Category
}
