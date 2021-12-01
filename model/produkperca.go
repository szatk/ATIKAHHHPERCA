package model

import (
	"time"

	"gorm.io/gorm"
)

type ProdukPerca struct {
	Id            uint           `gorm:"primaryKey; unique; not null" json:"id"`
	PercaSepatuId uint           `json:"PercaSepatuId"`
	PercaSepatu   PercaSepatu    `gorm:"foreignKey:PercaSepatuId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PercaBajuId   uint           `json:"PercaBajuId"`
	PercaBaju     PercaBaju      `gorm:"foreignKey:PercaBajuId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
