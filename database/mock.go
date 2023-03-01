package database

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
	"time"
)

type Mock struct {
	Error error
}

var mockedTime = time.Date(2023, 03, 1, 18, 20, 0, 0, time.UTC)

var mockedDatabase = map[string]entities.Currency{
	"BRL": {Code: "BRL", Name: "Real Brasileiro", USDConversionRate: 0.2, UpdatedAt: mockedTime},
	"EUR": {Code: "EUR", Name: "Euro", USDConversionRate: 1.05, UpdatedAt: mockedTime},
}

func (m Mock) Get(ctx context.Context, code string) (entities.Currency, error) {
	if currency, ok := mockedDatabase[code]; ok {
		return currency, nil
	}

	return entities.Currency{}, m.Error
}

func (m Mock) Create(ctx context.Context, currency entities.Currency) (interface{}, error) {
	return currency.Code, m.Error
}

func (m Mock) UpInsert(ctx context.Context, currency entities.Currency) error {
	return m.Error
}

func (m Mock) Update(ctx context.Context, code string, currency entities.Currency) error {
	return m.Error
}

func (m Mock) Delete(ctx context.Context, code string) error {
	return m.Error
}
