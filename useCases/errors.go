package useCases

import "github.com/pkg/errors"

var (
	ErrFromRequired  = errors.New("query param \"from\" is required")
	ErrToRequired    = errors.New("query param \"to\" is required")
	ErrInvalidAmount = errors.New("\"amount\" must be a valid float")

	ErrRequestTimeout = errors.New("request timeout")
)
