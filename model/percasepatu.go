package model

import (
	"time"

	"gorm.io/gorm"
)

type PercaSepatu struct {
	Id         uint           `gorm:"primaryKey; unique; not null" json:"id"`
	NamaProduk string         `json:"namaProduk"`
	Bahan      string         `json:"bahan"`
	Harga      string         `json:"harga"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
