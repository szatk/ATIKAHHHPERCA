package model

import (
	"time"

	"gorm.io/gorm"
)

type PekerjaPerca struct {
	Id             uint           `gorm:"primaryKey; unique; not null" json:"id"`
	HousePercaId   uint           `json:"housePercaId"`
	Nama           string         `gorm:"size:255; not null" json:"namaDepan"`
	NoTelepon      string         `gorm:"size:255;not null; unique" json:"noTelepon"`
	Alamat         string         `gorm:"size:1000;not null" json:"alamat"`
	Kabupaten_Kota string         `gorm:"size:255;not null" json:"kabupaten_kota"`
	Provinsi       string         `gorm:"size:255;not null" json:"provinsi"`
	HousePerca     HousePerca     `gorm:"foreignKey:HousePercaId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
