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
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	amount, _ := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)

	resp, err := c.UseCase.Convert(amount, from, to)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}
