package entities

import "context"

type QuotationData struct {
	Code      string `json:"code"`
	CodeIn    string `json:"codein"`
	Name      string `json:"name"`
	Ask       string `json:"ask"`
	UpdatedAt string `json:"timestamp"`
}

type QuotationRepository interface {
	GetCurrentUSDQuotation(ctx context.Context, code string) (*QuotationData, error)
}
