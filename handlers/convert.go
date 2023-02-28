package handlers

import (
	"github.com/lumacielz/challenge-bravo/presenters"
	"github.com/lumacielz/challenge-bravo/useCases"
	"net/http"
	"strconv"
)

type CurrencyController struct {
	UseCase   useCases.CurrencyUseCase
	Presenter presenters.CurrencyOutput
}

func (c CurrencyController) GetConversionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amount, _ := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)

	resp, err := c.UseCase.Convert(ctx, amount, from, to)

	if err != nil {
		c.Presenter.WriteError(w, err, 500)
	}

	c.Presenter.WriteResponse(w, resp)
}
