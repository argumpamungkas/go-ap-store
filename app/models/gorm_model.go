package models

import "time"

// struct reusable
type GormModel struct {
	ID          string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	CreatedDate time.Time
	UpdatedDate time.Time
}
