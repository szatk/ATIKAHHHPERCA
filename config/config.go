package config

import (
	"ATIKAH-PERCA/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "root:567890@tcp(127.0.0.1:3306)/perca?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initMigration()
}

var (
	DB *gorm.DB
)

func initMigration() {
	DB.AutoMigrate(
		&model.Admins{},
		// &model.Akreditasi{},
		// &model.Deskripsi{},
		// &model.Fakultas{},
		// &model.FasilitasJurusan{},
		// &model.FasilitasKampus{},
		// &model.FasilitasKelas{},
		// &model.Jurusans{},
		// &model.UKM{},
		// &model.Universitas{},
	)
}