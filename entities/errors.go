package entities

import (
	"errors"
	"fmt"
)

var (
	ErrCurrencyNotFound   = errors.New("Currency does not exist")
	ErrCodeRequired       = errors.New("Field code is required")
	ErrZeroConversionRate = errors.New("USDConversionRate must be greater than 0")

	ErrQuotationAPITimeout = errors.New("QuotationAPI took too long to answer")
)

func ErrUnexpected(status string) error {
	message := fmt.Sprintf("QuotationAPI returned an unexpected status code: %s", status)
	return errors.New(message)
}
