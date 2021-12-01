package main

import (
	"ATIKAH-PERCA/config"
	"ATIKAH-PERCA/route"
)

//MAIN FUNCTION
func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8080")
}
