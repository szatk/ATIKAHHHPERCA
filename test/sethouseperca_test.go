package test

import (
	"ATIKAH-PERCA/config"
	"ATIKAH-PERCA/route"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo/v4"
)

var (
	echoHandler *echo.Echo
	server      *httptest.Server
)

func init() {
	config.InitDB()
	config.DB.AutoMigrate()

	echoHandler = route.New()
	server = httptest.NewServer(echoHandler)

}

func TearUp() func() {
	return func() {
		config.InitDB()
		config.DB.Exec("TRUNCATE TABLE house_percas;")
	}
}

// GetHttpExpect Get cached expect, create new if nil
func GetHTTPExpect(t *testing.T) *httpexpect.Expect {
	if server == nil {
		server = httptest.NewServer(echoHandler)
	}
	return NewHTTPExpect(t)
}

func NewHTTPExpect(t *testing.T) *httpexpect.Expect {
	// TODO : printer set, only if the testing is failed
	// https://github.com/gavv/httpexpect/issues/69
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: server.URL,
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
		Reporter: t,
	})
}
