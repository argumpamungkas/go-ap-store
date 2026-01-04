package models

import "time"

// struct reusable
type GormModel struct {
	ID          string     `gorm:"size:36;not null;uniqueIndex;primary_key"`
	CreatedDate *time.Time `gorm:"autoCreateTime"`
	UpdatedDate *time.Time `gorm:"autoUpdateTime"`
}
