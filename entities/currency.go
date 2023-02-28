package entities

import (
	"context"
	"time"
)

type Currency struct {
	Code              string    `bson:"code"`
	Name              string    `bson:"name"`
	USDConversionRate float64   `bson:"USDConversionRate"`
	UpdatedAt         time.Time `bson:"updatedAt"`
}

type CurrencyRepository interface {
	Get(ctx context.Context, code string) (Currency, error)
	Create(ctx context.Context, currency Currency) error
	UpInsert(ctx context.Context, currency Currency) error
	Update(ctx context.Context, code string, currency Currency) error
	Delete(ctx context.Context, code string) error
}
