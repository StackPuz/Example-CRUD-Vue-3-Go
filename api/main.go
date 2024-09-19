package main

import (
	"app/config"
	"app/router"
)

func main() {
	config.SetupDatabase()
	router.SetupRouter()
}
