package dao

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Israel-Ferreira/hurb-currency/dto"
	"github.com/Israel-Ferreira/hurb-currency/exceptions"
	"github.com/Israel-Ferreira/hurb-currency/models"
)

type CurrencyRepo struct {
	Db *sql.DB
}

func (c CurrencyRepo) AddCurrencyPair(currencyDTO dto.CurrencyDTO) error {

	if currencyDTO.FromCoin == "" {
		return exceptions.ErrFromCoinEmptyOrNil
	}

	if currencyDTO.ToCoin == "" {
		return exceptions.ErrToCoinEmptyOrNil
	}

	fromExmplValue, err := strconv.ParseFloat(currencyDTO.ExampleFromValue, 64)

	if err != nil {
		return err
	}


	toExmplValue, err := strconv.ParseFloat(currencyDTO.ExampleToValue, 64)

	if err !=  nil {
		return err
	}

	
	fmt.Printf("FromValue: %.2f, ToValue: %.2f \n", fromExmplValue, toExmplValue)


	return nil
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
