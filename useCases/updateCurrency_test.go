package useCases

import (
	"context"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyUseCase_UpdateCurrency(t *testing.T) {
	type args struct {
		currencyRepositoryError error
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
			name: "error on database",
			args: args{
				currencyRepositoryError: entities.ErrCurrencyNotFound,
				currency: CurrencyRequest{
					Name:              "MyCoin",
					USDConversionRate: 2.0,
				},
			},
			wantErr: entities.ErrCurrencyNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrencyUseCase{
				CurrencyRepository: database.Mock{Error: tt.args.currencyRepositoryError},
			}
			err := c.UpdateCurrency(context.Background(), "", tt.args.currency)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
