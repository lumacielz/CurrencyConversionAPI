package database

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
)

type Mock struct {
	Error error
}

func (m Mock) Get(ctx context.Context, code string) (entities.Currency, error) {
	return entities.Currency{Code: code}, m.Error
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
