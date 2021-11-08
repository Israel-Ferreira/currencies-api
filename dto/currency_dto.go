package dto

type CurrencyDTO struct {
	FromCoin         string `json:"from"`
	ToCoin           string `json:"to"`
	ExampleFromValue string `json:"exampleFromValue"`
	ExampleToValue   string `json:"exampleToValue"`
}

func NewCurrencyDTOWithValues(from, to, exampleFromValue, exampleToValue string) (CurrencyDTO) {
	return CurrencyDTO{
		FromCoin:  from,
		ToCoin: to,
		ExampleFromValue: exampleFromValue,
		ExampleToValue: exampleToValue,
	}
}
