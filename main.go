package main

import (
	"mygram/config"
	"mygram/routers"
)

func main() {
	config.InitDB()

	routers.StartServer().Run(":9999")
}
