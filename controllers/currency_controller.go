package controllers

import (
	"fmt"

	"github.com/Israel-Ferreira/hurb-currency/config"
	"github.com/Israel-Ferreira/hurb-currency/dao"
	"github.com/gin-gonic/gin"
)

type CurrencyController struct {
	CurrencyRepo dao.CurrencyRepo
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

	cc :=  CurrencyController{CurrencyRepo: cr}

	currencyRoutes := r.Group("/currency")
	{
		currencyRoutes.GET("/", cc.GetConversionValue)
	}
}
