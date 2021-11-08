package dao

import (
	"database/sql"

	"github.com/Israel-Ferreira/hurb-currency/models"
)

type CurrencyRepo struct {
	Db *sql.DB
}

func (c CurrencyRepo) FindByFromCurrencyAndToCurrency(fromCt, toCt string) (models.Currency, error) {

	var currency models.Currency

	res, err := c.Db.Query(
		"select * from  CoinCurrencies where FromCoin = ? and ToCoin = ? ",
		fromCt,
		toCt,
	)

	if err != nil {
		return models.NewCurrency(), err
	}

	defer res.Close()

	if res.Next() {

		if err = res.Scan(&currency.From, &currency.To, &currency.Value); err != nil {
			return models.NewCurrency(), err
		}
	}

	return currency, nil

}
