package model

import (
	"time"

	"gorm.io/gorm"
)

type JenisPerca struct {
	Id        uint           `gorm:"primaryKey; unique; not null" json:"id"`
	NamaJenis string         `gorm:"size:255; not null" json:"namaJenis"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
