package useCases

import (
	"context"
	"errors"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyUseCase_DeleteCurrency(t *testing.T) {
	type args struct {
		currencyRepositoryError error
		code                    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success",
			args: args{
				code: "BRL",
			},
			wantErr: nil,
		},
		{
			name: "no found",
			args: args{
				code: "not found",
			},
			wantErr: entities.ErrCurrencyNotFound,
		},
		{
			name: "error deleting",
			args: args{
				code:                    "BRL",
				currencyRepositoryError: errors.New("error deleting currency"),
			},
			wantErr: errors.New("error deleting currency"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrencyUseCase{
				CurrencyRepository: database.Mock{Error: tt.args.currencyRepositoryError},
			}
			err := c.DeleteCurrency(context.Background(), tt.args.code)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
