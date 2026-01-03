package models

type Address struct {
	GormModel  `gorm:"embedded"`
	UserID     string `gorm:"size:36;index"` // foreign key
	User       User   // belongs to
	Name       string `gorm:"size:100;not null"`
	IsPrimary  bool   // user bisa menentukan alamat utama
	CityID     string `gorm:"size:100;"` // untuk menentukan ongkos kirim
	ProvinceID string `gorm:"size:100;"` // untuk menentukan ongkos kirim
	Address1   string `gorm:"size:255;"`
	Address2   string `gorm:"size:255;"`
	Phone      string `gorm:"size:100;"`
	Email      string `gorm:"size:100;"`
	PostCode   string `gorm:"size:100;"`
}
