package handlers

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/lumacielz/challenge-bravo/presenters"
	"github.com/lumacielz/challenge-bravo/useCases"
	"net/http"
	"strconv"
	"time"
)

type CurrencyController struct {
	UseCase         useCases.CurrencyUseCase
	OutputPresenter presenters.CurrencyOutput
	InputPresenter  presenters.CurrencyInput
	Timeout         time.Duration
}

func (c CurrencyController) GetConversionHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), c.Timeout)
	defer cancel()

	respC := make(chan *useCases.CurrencyConversionResponse)
	errC := make(chan error, 1)
	go func() {
		defer close(errC)
		defer close(respC)
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")
		amount, err := strconv.ParseFloat(r.URL.Query().Get("amount"), 64)
		if err != nil {
			errC <- useCases.ErrInvalidAmount
			return
		}

		if from == "" {
			errC <- useCases.ErrFromRequired
			return
		}

		if to == "" {
			errC <- useCases.ErrToRequired
			return
		}
		resp, err := c.UseCase.Convert(ctx, amount, from, to)
		if err != nil {
			errC <- err
			return
		}
		respC <- resp
	}()

	var status int
	select {
	case <-ctx.Done():
		c.OutputPresenter.WriteError(w, useCases.ErrRequestTimeout, http.StatusRequestTimeout)
		return
	case resp := <-respC:
		c.OutputPresenter.WriteResponse(w, &resp, http.StatusOK)
		return
	case err := <-errC:
		switch err {
		case useCases.ErrFromRequired, useCases.ErrToRequired, useCases.ErrInvalidAmount:
			status = http.StatusBadRequest
		case entities.ErrCurrencyNotFound:
			status = http.StatusNotFound
		default:
			status = http.StatusInternalServerError
		}
		c.OutputPresenter.WriteError(w, err, status)
	}
}
