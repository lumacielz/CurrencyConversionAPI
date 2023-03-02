package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/lumacielz/challenge-bravo/useCases"
	"net/http"
	"time"
)

func (c CurrencyController) DeleteCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 500*time.Millisecond)
	defer cancel()

	errC := make(chan error, 1)
	go func() {
		defer close(errC)

		code := chi.URLParam(r, "code")

		err := c.UseCase.DeleteCurrency(ctx, code)
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
		default:
			status = http.StatusInternalServerError
		}
		c.OutputPresenter.WriteError(w, err, status)
	}
}
