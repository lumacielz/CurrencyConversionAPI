package external

import (
	"context"
	"fmt"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestQuotationClient_GetCurrentUSDQuotation(t *testing.T) {
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	type args struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.QuotationData
		wantErr error
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				code: "BRL",
			},
			want: &entities.QuotationData{
				Code:      "BRL",
				CodeIn:    "USD",
				Name:      "Real Brasileiro/Dólar Americano",
				Ask:       "0.191",
				UpdatedAt: "1677636376",
			},
			wantErr: nil,
		},
		{
			name: "success with empty array",
			args: args{
				ctx:  context.Background(),
				code: "empty",
			},
			want:    nil,
			wantErr: entities.ErrCurrencyNotFound,
		},
		{
			name: "timeout",
			args: args{
				ctx: cancelledCtx,
			},
			want:    nil,
			wantErr: entities.ErrQuotationAPITimeout,
		},
		{
			name: "error not found",
			args: args{
				ctx:  context.Background(),
				code: "notFound",
			},
			want:    nil,
			wantErr: entities.ErrCurrencyNotFound,
		},
		{
			name: "unexpected error",
			args: args{
				ctx:  context.Background(),
				code: "error",
			},
			want:    nil,
			wantErr: entities.ErrUnexpected(fmt.Sprintf("%d %s", 500, http.StatusText(500))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := QuotationClient{
				Client:  http.DefaultClient,
				Url:     fmt.Sprintf("%s/json/%s", MockedServer.URL, "%s-USD"),
				Timeout: 5 * time.Second,
			}
			got, err := c.GetCurrentUSDQuotation(tt.args.ctx, tt.args.code)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
