package handlers

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/lumacielz/challenge-bravo/useCases"
	"io/ioutil"
	"net/http"
	"time"
)

func (c CurrencyController) NewCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Millisecond)
	defer cancel()

	respC := make(chan *useCases.NewCurrencyResponse)
	errC := make(chan error, 1)
	go func() {
		defer close(errC)
		defer close(respC)

		body, _ := ioutil.ReadAll(r.Body)
		currencyReq, err := c.InputPresenter.ParseRequest(body)
		if err != nil {
			errC <- err
			return
		}

		resp, err := c.UseCase.NewCurrency(ctx, currencyReq)
		if err != nil {
			errC <- err
			return
		}
		respC <- resp
	}()

	var status int
	select {
	case <-ctx.Done():
		c.OutputPresenter.WriteError(w, useCases.ErrRequestTimeout, http.StatusRequestTimeout)
		return
	case resp := <-respC:
		c.OutputPresenter.WriteResponse(w, &resp, http.StatusCreated)
		return
	case err := <-errC:
		switch err {
		case entities.ErrCodeRequired, entities.ErrZeroConversionRate:
			status = http.StatusUnprocessableEntity
		default:
			status = http.StatusInternalServerError
		}
		c.OutputPresenter.WriteError(w, err, status)
	}
}
