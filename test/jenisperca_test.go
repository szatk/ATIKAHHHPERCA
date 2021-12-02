package test

import (
	"ATIKAH-PERCA/config"
	m_jenisperca "ATIKAH-PERCA/model"
	"fmt"
	"net/http"
	"testing"
)

func init() {
	config.InitDB()

	jenisperca := &m_jenisperca.JenisPerca{
		NamaJenis: "Satin",
	}
	if err := config.DB.Create(&jenisperca).Error; err != nil {
		fmt.Println("error")
	}
}

func TestGetJenisPerca(t *testing.T) {
	HahaDown := HahaUp()
	defer HahaDown()

	e := GetHTTPExpect(t)

	result := e.GET("/jeniserca").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("message").String().Contains("hope")
	result.Value("data").Array().NotEmpty()
}
