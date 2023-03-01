package useCases

import (
	"context"
	"fmt"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/lumacielz/challenge-bravo/external"
	"github.com/stretchr/testify/assert"

	"net/http"
	"testing"
)

func TestCurrencyUseCase_UpdateCurrencyData(t *testing.T) {
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
			wantErr: entities.ErrUnexpected(fmt.Sprintf("%d %s", 500, http.StatusText(500))),
		},
		{
			name: "error on database",
			args: args{
				ctx:                     context.Background(),
				currencyRepositoryError: entities.ErrCurrencyNotFound,
				code:                    "ok",
			},
			wantErr: entities.ErrCurrencyNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrencyUseCase{
				CurrencyRepository: database.Mock{tt.args.currencyRepositoryError},
				QuotationRepository: external.QuotationClient{
					Client: http.DefaultClient,
					Url:    fmt.Sprintf("%s/json/%s", external.MockedServer.URL, "%s-USD"),
				},
			}
			err := c.UpdateCurrencyData(tt.args.ctx, tt.args.code)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
