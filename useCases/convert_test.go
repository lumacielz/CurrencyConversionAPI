package useCases

import (
	"context"
	"errors"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/external"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCurrencyUseCase_UpdateCurrencyData(t *testing.T) {
	server := httptest.NewServer(external.QuotationMock{})
	type args struct {
		ctx                     context.Context
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
				ctx:                     context.Background(),
				currencyRepositoryError: nil,
				code:                    "ok",
			},
			wantErr: nil,
		},
		{
			name: "error calling external",
			args: args{
				ctx:                     context.Background(),
				currencyRepositoryError: nil,
				code:                    "error",
			},
			wantErr: errors.New("erro"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrencyUseCase{
				CurrencyRepository: database.Mock{tt.args.currencyRepositoryError},
				QuotationClient:    external.QuotationClient{Client: http.DefaultClient, Url: server.URL + "/json/%s-USD"},
			}
			err := c.UpdateCurrencyData(tt.args.ctx, tt.args.code)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
