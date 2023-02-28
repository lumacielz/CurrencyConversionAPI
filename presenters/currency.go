package presenters

import (
	"encoding/json"
	"github.com/lumacielz/challenge-bravo/useCases"
	"net/http"
)

type CurrencyOutput interface {
	WriteResponse(w http.ResponseWriter, response useCases.CurrencyResponse) error
	WriteError(w http.ResponseWriter, err error, status int)
}

type JsonPresenter struct{}

type JsonFormatError struct {
	Error string `json:"error"`
}

func (p JsonPresenter) WriteResponse(w http.ResponseWriter, response useCases.CurrencyResponse) error {
	respJson, err := json.Marshal(response)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
	return nil
}

func (p JsonPresenter) WriteError(w http.ResponseWriter, err error, status int) {
	formatedError := JsonFormatError{Error: err.Error()}
	errJson, _ := json.Marshal(formatedError)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(errJson)
}
