package useCases

import (
	"github.com/lumacielz/challenge-bravo/entities"
	"time"
)

type CurrencyUseCase struct {
	Now                 func() time.Time
	CurrencyRepository  entities.CurrencyRepository
	QuotationRepository entities.QuotationRepository
}

type CurrencyConversionResponse struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type NewCurrencyResponse struct {
	Id interface{} `json:"_id"`
}

type CurrencyRequest struct {
	Code              string  `json:"code"`
	Name              string  `json:"name"`
	USDConversionRate float64 `json:"USDConversionRate"`
}

func validateCurrency(currency CurrencyRequest) error {
	if currency.Code == "" {
		return entities.ErrCodeRequired
	}
	if currency.USDConversionRate <= 0 {
		return entities.ErrZeroConversionRate
	}
	return nil
}

func validateUpdateRequest(request CurrencyRequest) error {
	if request.USDConversionRate <= 0 {
		return entities.ErrZeroConversionRate
	}
	return nil
}
