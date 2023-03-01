package useCases

import (
	"context"
	"errors"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/lumacielz/challenge-bravo/external"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCurrencyUseCase_UpdateCurrencyData(t *testing.T) {
	type args struct {
		currencyRepositoryError  error
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
				currencyRepositoryError:  nil,
				quotationRepositoryError: nil,
			},
			wantErr: nil,
		},
		{
			name: "error calling external",
			args: args{
				currencyRepositoryError:  nil,
				quotationRepositoryError: entities.ErrUnexpected("500 internal server"),
			},
			wantErr: entities.ErrUnexpected("500 internal server"),
		},
		{
			name: "error on database",
			args: args{
				currencyRepositoryError:  entities.ErrCurrencyNotFound,
				quotationRepositoryError: entities.ErrCurrencyNotFound,
			},
			wantErr: entities.ErrCurrencyNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrencyUseCase{
				CurrencyRepository:  database.Mock{Error: tt.args.currencyRepositoryError},
				QuotationRepository: external.QuotationMock{Error: tt.args.quotationRepositoryError},
			}
			err := c.UpdateCurrencyData(context.Background(), "")
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestCurrencyUseCase_Convert(t *testing.T) {
	type args struct {
		mockedTime               time.Time
		currencyRepositoryError  error
		quotationRepositoryError error
		amount                   float64
		from                     string
		to                       string
	}
	tests := []struct {
		name    string
		args    args
		want    CurrencyConversionResponse
		wantErr error
	}{
		{
			name: "success with data up to date",
			args: args{
				mockedTime:              time.Date(2023, 03, 1, 18, 20, 10, 0, time.UTC),
				currencyRepositoryError: nil,
				amount:                  50,
				from:                    "BRL",
				to:                      "EUR",
			},
			want:    CurrencyConversionResponse{Value: "9.524", Currency: "EUR"},
			wantErr: nil,
		},
		{
			name: "success updating data",
			args: args{
				mockedTime:               time.Date(2023, 03, 1, 18, 31, 10, 0, time.UTC),
				currencyRepositoryError:  nil,
				quotationRepositoryError: nil,
				amount:                   50,
				from:                     "BRL",
				to:                       "EUR",
			},
			want:    CurrencyConversionResponse{Value: "9.524", Currency: "EUR"},
			wantErr: nil,
		},
		{
			name: "error updating data - ignore and use old data",
			args: args{
				mockedTime:               time.Date(2023, 03, 1, 18, 31, 10, 0, time.UTC),
				currencyRepositoryError:  nil,
				quotationRepositoryError: errors.New("error"),
				amount:                   50,
				from:                     "BRL",
				to:                       "EUR",
			},
			want:    CurrencyConversionResponse{Value: "9.524", Currency: "EUR"},
			wantErr: nil,
		},
		{
			name: "error getting from database",
			args: args{
				mockedTime:               time.Date(2023, 03, 1, 18, 31, 10, 0, time.UTC),
				currencyRepositoryError:  entities.ErrCurrencyNotFound,
				quotationRepositoryError: nil,
				amount:                   50,
				from:                     "BL",
				to:                       "EUR",
			},
			want:    CurrencyConversionResponse{},
			wantErr: entities.ErrCurrencyNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrencyUseCase{
				Now:                 func() time.Time { return tt.args.mockedTime },
				CurrencyRepository:  database.Mock{Error: tt.args.currencyRepositoryError},
				QuotationRepository: external.QuotationMock{Error: tt.args.quotationRepositoryError},
			}
			got, err := c.Convert(context.Background(), tt.args.amount, tt.args.from, tt.args.to)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
