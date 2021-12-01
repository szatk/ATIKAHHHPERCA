package model

import (
	"time"

	"gorm.io/gorm"
)

type InfoHouse struct {
	Id             uint           `gorm:"primaryKey; unique; not null" json:"id"`
	HousePercaId   uint           `json:"housePercaId"`
	HousePerca     HousePerca     `gorm:"foreignKey:HousePercaId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Pelatihan      string         `json:"pelatihan"`
	HargaPelatihan string         `json:"hargapelatihan"`
	Manfaat        string         `json:"manfaat"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
