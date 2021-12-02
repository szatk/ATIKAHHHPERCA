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
	echoHandler1 *echo.Echo
	server1      *httptest.Server
)

func init() {
	config.InitDB()
	config.DB.AutoMigrate()

	echoHandler1 = route.New()
	server1 = httptest.NewServer(echoHandler1)

}

func HahaUp() func() {
	return func() {
		config.InitDB()
		config.DB.Exec("TRUNCATE TABLE jenis_percas;")
	}
}

// GetHttpExpect Get cached expect, create new if nil
func GetHTTPExpected(t *testing.T) *httpexpect.Expect {
	if server1 == nil {
		server1 = httptest.NewServer(echoHandler1)
	}
	return NewHTTPExpect(t)
}

func NewHTTPExpected(t *testing.T) *httpexpect.Expect {
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
