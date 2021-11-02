package models

type Currency struct {
	Pair string
	From string
	To string
	Value float64
}


func NewCurrency() Currency {
	return Currency{}
}

