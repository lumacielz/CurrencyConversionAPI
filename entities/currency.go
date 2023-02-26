package entities

import "context"

type Currency struct {
	Code              string
	Name              string
	USDConversionRate float64
}

type CurrencyRepository interface {
	Get(ctx context.Context, code string) (Currency, error)
	//Create(currency Currency) error
	//Update(currency Currency) error
	//Delete(code string) error
}
