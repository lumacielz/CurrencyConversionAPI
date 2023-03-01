package useCases

import (
	"github.com/lumacielz/challenge-bravo/entities"
)

type CurrencyUseCase struct {
	CurrencyRepository  entities.CurrencyRepository
	QuotationRepository entities.QuotationRepository
}

type CurrencyConversionResponse struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

type NewCurrencyResponse struct {
	Id interface{} `json:"_id"`
}

type CurrencyRequest struct {
	Code              string  `json:"code"`
	Name              string  `json:"name"`
	USDConversionRate float64 `json:"USDConversionRate"`
}
