package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	GormModel           `gorm:"embedded"`
	UserID              string `gorm:"size:36;index;"`
	User                User
	OrderItems          []OrderItem
	OrderCustomers      *OrderCustomer
	Code                string `gorm:"size:50;index"`
	Status              int
	OrderDate           time.Time
	PaymentDue          time.Time
	PaymentStatus       string          `gorm:"size:50;index;"`
	PaymentToken        string          `gorm:"size:100;index;"`
	BaseTotalPrice      decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxAmount           decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxPercent          decimal.Decimal `gorm:"type:decimal(10,2)"`
	DiscountAmount      decimal.Decimal `gorm:"type:decimal(16,2)"`
	DiscountPercent     decimal.Decimal `gorm:"type:decimal(10,2)"`
	ShippingCost        decimal.Decimal `grom:"type:decimal(16,2)"`
	GrandTotal          decimal.Decimal `grom:"type:decimal(16,2)"`
	Note                string          `gorm:"type:text"`
	ShippingCourier     string          `gorm:"size:100"`
	ShippingServiceName string          `gorm:"size:100"`
	ApprovedBy          string          `gorm:"size:36"`
	ApporvedDate        time.Time
	CancelledBy         string `gorm:"size:36"`
	CancelledDate       time.Time
	CancellationNote    string `gorm:"size:255"`
	DeletedDate         gorm.DeletedAt
}
