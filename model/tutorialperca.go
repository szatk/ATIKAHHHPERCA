package model

import (
	"time"

	"gorm.io/gorm"
)

type TutorialPerca struct {
	Id        uint           `gorm:"primaryKey; unique; not null" json:"id"`
	Bahan     string         `gorm:"size:255; not null" json:"bahan"`
	Alat      string         `gorm:"size:255; not null" json:"alat"`
	Cara1     string         `gorm:"size:1000; not null" json:"cara1"`
	Cara2     string         `gorm:"size:1000; not null" json:"cara2"`
	Cara3     string         `gorm:"size:1000; not null" json:"cara3"`
	Cara4     string         `gorm:"size:1000; not null" json:"cara4"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
