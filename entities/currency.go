package entities

import (
	"context"
	"time"
)

type Currency struct {
	Code              string
	Name              string
	USDConversionRate float64
	UpdatedAt         time.Time
}

type CurrencyRepository interface {
	Get(ctx context.Context, code string) (Currency, error)
	//Create(currency Currency) error
	//Update(currency Currency) error
	//Delete(code string) error
}
