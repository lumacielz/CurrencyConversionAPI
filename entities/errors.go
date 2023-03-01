package entities

import (
	"errors"
	"fmt"
)

var ErrCurrencyNotFound = errors.New("Currency does not exist")

func ErrUnexpected(status string) error {
	message := fmt.Sprintf("QuotationAPI returned an unexpected status code: %s", status)
	return errors.New(message)
}
