package handlers

import (
	"io/ioutil"
	"net/http"
)

func (c CurrencyController) NewCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)
	currencyReq, err := c.InputPresenter.Parse(body)
	insertedId, err := c.UseCase.NewCurrency(ctx, currencyReq)
	if err != nil {
		c.OutputPresenter.WriteError(w, err, 500)
		return
	}

	c.OutputPresenter.WriteResponse(w, insertedId, http.StatusCreated)
}
