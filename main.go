package main

import (
	"fmt"

	"github.com/Israel-Ferreira/hurb-currency/config"
	"github.com/Israel-Ferreira/hurb-currency/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Teste")

	database := "currencies"
	url := "localhost:35060"
	username := "root"
	password := "root"

	config.OpenConnection(url, database, username, password)

	defer config.DB.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello World",
		})
	})

	controllers.LoadCurrencyRoutes(r)

	r.Run()
}
