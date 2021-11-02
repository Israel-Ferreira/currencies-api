package controllers

import "github.com/gin-gonic/gin"


type CurrencyController struct {
	
}



func LoadCurrencyRoutes(r *gin.Engine) {
	currencyRoutes := r.Group("/currency")
	{
		currencyRoutes.GET("/", nil)
	}
}