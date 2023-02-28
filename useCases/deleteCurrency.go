package useCases

import "context"

func (c CurrencyUseCase) DeleteCurrency(ctx context.Context, code string) error {
	return c.CurrencyRepository.Delete(ctx, code)
}
