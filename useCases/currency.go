package useCases

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
)

type CurrencyUseCase struct {
	CurrencyRepository entities.CurrencyRepository
}

//TODO mover para presenter
type CurrencyResponse struct {
	Value    float64
	Currency string
}

func (c CurrencyUseCase) Convert(amount float64, from, to string) (CurrencyResponse, error) {
	//TODO usar errgroup
	originCurrencyData, err := c.CurrencyRepository.Get(context.Background(), from)
	if err != nil {
		return CurrencyResponse{}, err
	}

	destinationCurrencyData, err := c.CurrencyRepository.Get(context.Background(), to)
	if err != nil {
		return CurrencyResponse{}, err
	}

	destinationValue := originCurrencyData.USDConversionRate / destinationCurrencyData.USDConversionRate * amount
	response := CurrencyResponse{
		Value:    destinationValue,
		Currency: to,
	}

	return response, nil
}
