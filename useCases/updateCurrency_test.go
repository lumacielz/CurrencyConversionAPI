package useCases

import (
	"context"
	"errors"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyUseCase_UpdateCurrency(t *testing.T) {
	type args struct {
		currencyRepositoryError error
		code                    string
		currency                CurrencyRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success",
			args: args{
				currencyRepositoryError: nil,
				code:                    "BRL",
				currency: CurrencyRequest{
					Name:              "MyCoin",
					USDConversionRate: 2.0,
				},
			},
			wantErr: nil,
		},
		{
			name: "validation error",
			args: args{
				currencyRepositoryError: nil,
				currency:                CurrencyRequest{},
			},
			wantErr: entities.ErrZeroConversionRate,
		},
		{
			name: "error not found",
			args: args{
				code: "nf",
				currency: CurrencyRequest{
					Name:              "MyCoin",
					USDConversionRate: 2.0,
				},
			},
			wantErr: entities.ErrCurrencyNotFound,
		},
		{
			name: "error on database",
			args: args{
				currencyRepositoryError: errors.New("error updating"),
				code:                    "BRL",
				currency: CurrencyRequest{
					Name:              "MyCoin",
					USDConversionRate: 2.0,
				},
			},
			wantErr: errors.New("error updating"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrencyUseCase{
				CurrencyRepository: database.Mock{Error: tt.args.currencyRepositoryError},
			}
			err := c.UpdateCurrency(context.Background(), tt.args.code, tt.args.currency)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
