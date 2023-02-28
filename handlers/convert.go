package handlers

import (
	"encoding/json"
	"github.com/lumacielz/challenge-bravo/useCases"
	"net/http"
	"strconv"
)

type CurrencyController struct {
	UseCase useCases.CurrencyUseCase
}

func (c CurrencyController) GetConversionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amount, _ := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)

	resp, err := c.UseCase.Convert(ctx, amount, from, to)

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	//TODO tratar erro
	respJson, err := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJson)
}
