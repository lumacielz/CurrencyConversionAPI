package useCases

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
)

func (c CurrencyUseCase) NewCurrency(ctx context.Context, currency CurrencyRequest) (*NewCurrencyResponse, error) {
	err := validateCurrency(currency)
	if err != nil {
		return nil, err
	}

	currencyEntity := entities.Currency{
		Code:              currency.Code,
		Name:              currency.Name,
		USDConversionRate: currency.USDConversionRate,
	}
	id, err := c.CurrencyRepository.Create(ctx, currencyEntity)
	return &NewCurrencyResponse{Id: id}, err
}
