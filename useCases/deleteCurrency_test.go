package useCases

import (
	"context"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyUseCase_DeleteCurrency(t *testing.T) {
	type args struct {
		currencyRepositoryError error
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
			},
			wantErr: nil,
		},
		{
			name: "error",
			args: args{
				currencyRepositoryError: entities.ErrCurrencyNotFound,
			},
			wantErr: entities.ErrCurrencyNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrencyUseCase{
				CurrencyRepository: database.Mock{Error: tt.args.currencyRepositoryError},
			}
			err := c.DeleteCurrency(context.Background(), "")
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
