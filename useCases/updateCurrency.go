package useCases

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
)

func (c CurrencyUseCase) UpdateCurrency(ctx context.Context, code string, currency CurrencyRequest) error {
	err := validateUpdateRequest(currency)
	if err != nil {
		return err
	}

	currencyEntity := entities.Currency{
		Name:              currency.Name,
		USDConversionRate: currency.USDConversionRate,
	}
	err = c.CurrencyRepository.Update(ctx, code, currencyEntity)
	return err
}
