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
