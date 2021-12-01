package model

import (
	"time"

	"gorm.io/gorm"
)

type HousePerca struct {
	Id             int            `gorm:"primaryKey; auto_increment; not null" json:"id"`
	NamaUsaha      string         `gorm:"size:255; not null" json:"namaUsaha"`
	NoTelepon      string         `gorm:"size:255; not null; unique" json:"noTelepon"`
	Email          string         `gorm:"size:255; not null; unique" json:"email"`
	Alamat         string         `gorm:"size:1000; not null" json:"alamat"`
	Kabupaten_Kota string         `gorm:"size:255; not null" json:"kabupaten_kota"`
	Provinsi       string         `gorm:"size:255; not null" json:"provinsi"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
