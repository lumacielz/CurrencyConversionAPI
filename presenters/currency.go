package presenters

import (
	"encoding/json"
	"github.com/lumacielz/challenge-bravo/useCases"
	"net/http"
)

type CurrencyOutput interface {
	WriteResponse(w http.ResponseWriter, response interface{}, status int) error
	WriteError(w http.ResponseWriter, err error, status int)
}

type CurrencyInput interface {
	ParseRequest(body []byte) (useCases.CurrencyRequest, error)
}

type JsonFormatError struct {
	Error string `json:"error"`
}

type JsonPresenter struct{}

func (p JsonPresenter) ParseRequest(body []byte) (useCases.CurrencyRequest, error) {
	var req useCases.CurrencyRequest
	err := json.Unmarshal(body, &req)
	return req, err
}

func (p JsonPresenter) WriteResponse(w http.ResponseWriter, response interface{}, status int) error {
	respJson, err := json.Marshal(response)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
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
