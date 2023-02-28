package handlers

import (
	"github.com/lumacielz/challenge-bravo/presenters"
	"github.com/lumacielz/challenge-bravo/useCases"
	"net/http"
	"strconv"
)

type CurrencyController struct {
	UseCase         useCases.CurrencyUseCase
	OutputPresenter presenters.CurrencyOutput
	InputPresenter  presenters.CurrencyInput
}

func (c CurrencyController) GetConversionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amount, _ := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)

	resp, err := c.UseCase.Convert(ctx, amount, from, to)

	if err != nil {
		c.OutputPresenter.WriteError(w, err, 500)
		return
	}

	c.OutputPresenter.WriteResponse(w, &resp, http.StatusOK)
}
