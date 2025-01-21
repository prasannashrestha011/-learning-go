package main

import (
	databaseconfig "main/cmd/configs/databaseConfig"
	"main/cmd/containers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	DB := databaseconfig.Connect()
	containers.InitUserContainer(r, DB)
	r.Run()

}
