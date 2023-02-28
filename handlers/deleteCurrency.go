package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (c CurrencyController) DeleteCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := chi.URLParam(r, "code")

	err := c.UseCase.CurrencyRepository.Delete(ctx, code)
	if err != nil {
		c.OutputPresenter.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	c.OutputPresenter.WriteResponse(w, nil, http.StatusNoContent)
}
