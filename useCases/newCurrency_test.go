package useCases

import (
	"context"
	"errors"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyUseCase_NewCurrency(t *testing.T) {
	type args struct {
		currencyRepositoryError error
		currency                CurrencyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    NewCurrencyResponse
		wantErr error
	}{
		{
			name: "success",
			args: args{
				currency: CurrencyRequest{
					Code:              "MyCoin",
					USDConversionRate: 2.0,
				},
			},
			want: NewCurrencyResponse{
				Id: "MyCoin",
			},
			wantErr: nil,
		},
		{
			name: "validation error",
			args: args{
				currency: CurrencyRequest{},
			},
			wantErr: entities.ErrCodeRequired,
		},
		{
			name: "error on database",
			args: args{
				currencyRepositoryError: errors.New("error inserting into database"),
				currency: CurrencyRequest{
					Code:              "MyCoin",
					USDConversionRate: 2.0,
				},
			},
			want:    NewCurrencyResponse{},
			wantErr: errors.New("error inserting into database"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrencyUseCase{
				CurrencyRepository: database.Mock{Error: tt.args.currencyRepositoryError},
			}
			got, err := c.NewCurrency(context.Background(), tt.args.currency)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
