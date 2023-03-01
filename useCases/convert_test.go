package useCases

import (
	"context"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/lumacielz/challenge-bravo/external"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrencyUseCase_UpdateCurrencyData(t *testing.T) {
	type args struct {
		ctx                      context.Context
		currencyRepositoryError  error
		quotationRepositoryResp  *entities.QuotationData
		quotationRepositoryError error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:                      context.Background(),
				currencyRepositoryError:  nil,
				quotationRepositoryResp:  &entities.QuotationData{Code: "OK"},
				quotationRepositoryError: nil,
			},
			wantErr: nil,
		},
		{
			name: "error calling external",
			args: args{
				ctx:                      context.Background(),
				currencyRepositoryError:  nil,
				quotationRepositoryError: entities.ErrUnexpected("500 internal server"),
			},
			wantErr: entities.ErrUnexpected("500 internal server"),
		},
		{
			name: "error on database",
			args: args{
				ctx:                      context.Background(),
				currencyRepositoryError:  entities.ErrCurrencyNotFound,
				quotationRepositoryError: entities.ErrCurrencyNotFound,
			},
			wantErr: entities.ErrCurrencyNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrencyUseCase{
				CurrencyRepository:  database.Mock{tt.args.currencyRepositoryError},
				QuotationRepository: external.QuotationMock{Resp: tt.args.quotationRepositoryResp, Err: tt.args.quotationRepositoryError},
			}
			err := c.UpdateCurrencyData(tt.args.ctx, "")
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
