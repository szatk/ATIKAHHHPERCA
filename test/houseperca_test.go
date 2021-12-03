package test

import (
	"ATIKAH-PERCA/config"
	m_houseperca "ATIKAH-PERCA/model"
	"fmt"
	"net/http"
	"testing"
)

func init() {
	config.InitDB()

	houseperca := &m_houseperca.HousePerca{
		NamaUsaha: "Test2",
	}
	if err := config.DB.Create(&houseperca).Error; err != nil {
		fmt.Println("error")
	}
}

func TestGetHousePerca(t *testing.T) {
	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.GET("/houseperca").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("message").String().Contains("Successful retrieve data")
	result.Value("data").Array().NotEmpty()
}
