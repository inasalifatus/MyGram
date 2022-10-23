package main

import (
	"mygram/config"
	"mygram/routers"
)

func main() {
	config.InitDB()
	config.GetDB()

	routers.StartServer().Run(":9999")
}
