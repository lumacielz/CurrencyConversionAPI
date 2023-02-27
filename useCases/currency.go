package useCases

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/lumacielz/challenge-bravo/external"
	"time"
)

type CurrencyUseCase struct {
	CurrencyRepository entities.CurrencyRepository
	QuotationClient    external.QuotationClient
}

//TODO mover para presenter
type CurrencyResponse struct {
	Value    float64
	Currency string
}

func (c *CurrencyUseCase) Convert(ctx context.Context, amount float64, from, to string) (CurrencyResponse, error) {
	//TODO usar errgroup
	originCurrencyData, err := c.CurrencyRepository.Get(ctx, from)
	if err != nil {
		return CurrencyResponse{}, err
	}

	destinationCurrencyData, err := c.CurrencyRepository.Get(ctx, to)
	if err != nil {
		return CurrencyResponse{}, err
	}

	c.UpdateCurrencyData(ctx, originCurrencyData)
	c.UpdateCurrencyData(ctx, destinationCurrencyData)

	destinationValue := originCurrencyData.USDConversionRate / destinationCurrencyData.USDConversionRate * amount
	response := CurrencyResponse{
		Value:    destinationValue,
		Currency: to,
	}

	return response, nil
}

func (c *CurrencyUseCase) UpdateCurrencyData(ctx context.Context, currencyData entities.Currency) {
	if time.Now().After(currencyData.UpdatedAt.Add(30 * time.Second)) {
		resp, _ := c.QuotationClient.GetCurrentUSDQuotation(ctx, currencyData.Code)
		if resp != nil {
			//TODO atualiza dados no banco
		}
	}
}
