package controllers

import (
	"fmt"
	"net/http"

	"github.com/Israel-Ferreira/hurb-currency/config"
	"github.com/Israel-Ferreira/hurb-currency/dao"
	"github.com/Israel-Ferreira/hurb-currency/dto"
	"github.com/gin-gonic/gin"
)

type CurrencyController struct {
	CurrencyRepo dao.CurrencyRepo
}

func (cc CurrencyController) AddCurrencyPair(c *gin.Context) {
	var currencyDTO dto.CurrencyDTO

	if err := c.ShouldBindJSON(&currencyDTO); err != nil {
		c.JSON(400, gin.H{
			"err": "Error to get sentet json",
		})

		return
	}

	if err := cc.CurrencyRepo.AddCurrencyPair(currencyDTO); err != nil {
		c.JSON(422, gin.H{
			"err": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": "currency added successful",
	})

}

func (cc CurrencyController) GetConversionValue(c *gin.Context) {
	from := c.Params.ByName("from")
	to := c.Params.ByName("to")

	currency, err := cc.CurrencyRepo.FindByFromCurrencyAndToCurrency(from, to)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(400, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(200, currency)
}

func LoadCurrencyRoutes(r *gin.Engine) {
	cr := dao.CurrencyRepo{Db: config.DB}

	cc := CurrencyController{CurrencyRepo: cr}

	currencyRoutes := r.Group("/currency")
	{
		currencyRoutes.GET("/", cc.GetConversionValue)
		currencyRoutes.POST("/", cc.AddCurrencyPair)
	}
}
