package config

// 6180063e-b6de-439b-ad7a-db3b7f93b0a5
import (
	"os"
	"fmt"

	"ATIKAH-PERCA/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)


var DB *gorm.DB

func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	DB_NAME := os.Getenv("DB_NAME")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// }

// func InitDB() {
// 	// dsn := "root:567890@tcp(127.0.0.1:3306)/perca?charset=utf8mb4&parseTime=True&loc=Local"
// 	dsn := "root:567890@tcp(host.docker.internal:3306)/perca?charset=utf8mb4&parseTime=True&loc=Local"
// 	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initMigration()
}

// var (
// 	DB *gorm.DB
// )

func initMigration() {
	DB.AutoMigrate(
		&model.Admins{},
		&model.HousePerca{},
		&model.InfoHouse{},
		&model.JenisPerca{},
		&model.PekerjaPerca{},
		&model.PercaBaju{},
		&model.PercaSepatu{},
		&model.TutorialPerca{},
	)
}
