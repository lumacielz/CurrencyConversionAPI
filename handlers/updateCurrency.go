package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/lumacielz/challenge-bravo/useCases"
	"io/ioutil"
	"net/http"
	"time"
)

func (c CurrencyController) UpdateCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Millisecond)
	defer cancel()

	errC := make(chan error, 1)
	go func() {
		defer close(errC)

		code := chi.URLParam(r, "code")

		body, _ := ioutil.ReadAll(r.Body)
		currencyReq, err := c.InputPresenter.ParseRequest(body)
		if err != nil {
			errC <- err
			return
		}

		err = c.UseCase.UpdateCurrency(ctx, code, currencyReq)
		errC <- err
	}()

	var status int
	select {
	case <-ctx.Done():
		c.OutputPresenter.WriteError(w, useCases.ErrRequestTimeout, http.StatusRequestTimeout)
		return
	case err := <-errC:
		switch err {
		case nil:
			c.OutputPresenter.WriteResponse(w, nil, http.StatusNoContent)
			return
		case entities.ErrCurrencyNotFound:
			status = http.StatusNotFound
		case entities.ErrZeroConversionRate:
			status = http.StatusUnprocessableEntity
		default:
			status = http.StatusInternalServerError
		}
		c.OutputPresenter.WriteError(w, err, status)
	}
}
