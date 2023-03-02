package handlers

import (
	"bytes"
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/external"
	"github.com/lumacielz/challenge-bravo/presenters"
	"github.com/lumacielz/challenge-bravo/useCases"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestCurrencyController_UpdateCurrencyHandler(t *testing.T) {
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	type args struct {
		ctx             context.Context
		code            string
		bodyPath        string
		repositoryError error
		wantStatus      int
		wantPath        string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				ctx:        context.Background(),
				code:       "BRL",
				bodyPath:   "/requests/newCurrency.json",
				wantStatus: http.StatusNoContent,
				wantPath:   "/responses/empty.json",
			},
		},
		{
			name: "json unmarshal error",
			args: args{
				ctx:        context.Background(),
				bodyPath:   "/requests/empty.json",
				wantStatus: http.StatusInternalServerError,
				wantPath:   "/responses/errUnmarshal.json",
			},
		},
		{
			name: "timeout",
			args: args{
				ctx:        cancelledCtx,
				wantStatus: http.StatusRequestTimeout,
				wantPath:   "/responses/errTimeout.json",
			},
		},
		{
			name: "validation error - zero rate",
			args: args{
				ctx:        context.Background(),
				bodyPath:   "/requests/currencyWithZeroRate.json",
				wantStatus: http.StatusUnprocessableEntity,
				wantPath:   "/responses/errZeroRateUnprocessableEntity.json",
			},
		},
		{
			name: "currency not found",
			args: args{
				ctx:        context.Background(),
				code:       "BL",
				bodyPath:   "/requests/newCurrency.json",
				wantStatus: http.StatusNotFound,
				wantPath:   "/responses/errNotFound.json",
			},
		},
		{
			name: "internal server error",
			args: args{
				ctx:             context.Background(),
				code:            "BRL",
				bodyPath:        "/requests/newCurrency.json",
				repositoryError: errors.New("unexpected error"),
				wantStatus:      http.StatusInternalServerError,
				wantPath:        "/responses/errInternalServer.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CurrencyController{
				UseCase: useCases.CurrencyUseCase{
					Now:                 func() time.Time { return time.Now() },
					CurrencyRepository:  database.Mock{Error: tt.args.repositoryError},
					QuotationRepository: external.QuotationMock{},
				},
				OutputPresenter: presenters.JsonPresenter{},
				InputPresenter:  presenters.JsonPresenter{},
			}

			w := httptest.NewRecorder()

			root, _ := os.Getwd()
			body, _ := ioutil.ReadFile(root + tt.args.bodyPath)
			r := httptest.NewRequest("PUT", "/", bytes.NewReader(body))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("code", tt.args.code)
			r = r.WithContext(context.WithValue(tt.args.ctx, chi.RouteCtxKey, rctx))

			c.UpdateCurrencyHandler(w, r)

			respBody, _ := ioutil.ReadAll(w.Body)
			wantBody, _ := ioutil.ReadFile(root + tt.args.wantPath)
			assert.Equal(t, string(wantBody), string(respBody))
			assert.Equal(t, tt.args.wantStatus, w.Code)
		})
	}
}
