package handlers

import (
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
)

func (c CurrencyController) UpdateCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := chi.URLParam(r, "code")

	body, err := ioutil.ReadAll(r.Body)
	currencyReq, err := c.InputPresenter.ParseRequest(body)

	err = c.UseCase.UpdateCurrency(ctx, code, currencyReq)
	if err != nil {
		c.OutputPresenter.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	c.OutputPresenter.WriteResponse(w, nil, http.StatusNoContent)
}
