package handlers

import (
	"context"
	"errors"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/external"
	"github.com/lumacielz/challenge-bravo/presenters"
	"github.com/lumacielz/challenge-bravo/useCases"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestCurrencyController_GetConversionHandler(t *testing.T) {
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel()
	type args struct {
		ctx             context.Context
		from            string
		to              string
		amount          string
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
				from:       "BRL",
				to:         "EUR",
				amount:     "50.0",
				wantStatus: http.StatusOK,
				wantPath:   "/responses/convert.json",
			},
		},
		{
			name: "bad request from required",
			args: args{
				ctx:        context.Background(),
				from:       "",
				to:         "EUR",
				amount:     "50.0",
				wantStatus: http.StatusBadRequest,
				wantPath:   "/responses/errFromBadRequest.json",
			},
		},
		{
			name: "bad request to required",
			args: args{
				ctx:        context.Background(),
				from:       "BRL",
				to:         "",
				amount:     "50.0",
				wantStatus: http.StatusBadRequest,
				wantPath:   "/responses/errToBadRequest.json",
			},
		},
		{
			name: "timeout",
			args: args{
				ctx:        cancelledCtx,
				from:       "BRL",
				to:         "EUR",
				amount:     "50.0",
				wantStatus: http.StatusRequestTimeout,
				wantPath:   "/responses/errTimeout.json",
			},
		},
		{
			name: "not found",
			args: args{
				ctx:        context.Background(),
				from:       "BR",
				to:         "EUR",
				amount:     "50.0",
				wantStatus: http.StatusNotFound,
				wantPath:   "/responses/errNotFound.json",
			},
		},
		{
			name: "internal server error",
			args: args{
				ctx:             context.Background(),
				from:            "BR",
				to:              "EUR",
				amount:          "50.0",
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
					Now: func() time.Time {
						return time.Now()
					},
					CurrencyRepository:  database.Mock{Error: tt.args.repositoryError},
					QuotationRepository: external.QuotationMock{},
				},
				OutputPresenter: presenters.JsonPresenter{},
				InputPresenter:  presenters.JsonPresenter{},
			}

			w := httptest.NewRecorder()

			r := httptest.NewRequest("GET", "/", nil)
			r = r.WithContext(tt.args.ctx)

			q := url.Values{}
			q.Add("from", tt.args.from)
			q.Add("to", tt.args.to)
			q.Add("amount", tt.args.amount)
			r.URL.RawQuery = q.Encode()

			c.GetConversionHandler(w, r)

			body, _ := ioutil.ReadAll(w.Body)
			root, _ := os.Getwd()
			wantBody, _ := ioutil.ReadFile(root + tt.args.wantPath)
			assert.Equal(t, string(wantBody), string(body))
			assert.Equal(t, tt.args.wantStatus, w.Code)
		})
	}
}
