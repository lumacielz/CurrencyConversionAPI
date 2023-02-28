package useCases

import (
	"context"
	"fmt"
	"github.com/lumacielz/challenge-bravo/entities"
	"github.com/lumacielz/challenge-bravo/external"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"strings"
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
	if shouldUpdateCurrencyData(originCurrencyData, err) {
		c.UpdateCurrencyData(ctx, from)
		originCurrencyData, err = c.CurrencyRepository.Get(ctx, from)
	}
	if err != nil {
		return CurrencyResponse{}, err
	}

	destinationCurrencyData, err := c.CurrencyRepository.Get(ctx, to)
	if shouldUpdateCurrencyData(destinationCurrencyData, err) {
		c.UpdateCurrencyData(ctx, to)
		destinationCurrencyData, err = c.CurrencyRepository.Get(ctx, to)
	}

	if err != nil {
		return CurrencyResponse{}, err
	}

	destinationValue := originCurrencyData.USDConversionRate / destinationCurrencyData.USDConversionRate * amount
	response := CurrencyResponse{
		Value:    destinationValue,
		Currency: to,
	}

	return response, nil
}

func (c *CurrencyUseCase) UpdateCurrencyData(ctx context.Context, code string) error {
	resp, err := c.QuotationClient.GetCurrentUSDQuotation(ctx, code)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	rate, _ := strconv.ParseFloat(resp.Ask, 64)

	var name string
	if names := strings.Split(resp.Name, "/"); len(names) > 0 {
		name = names[0]
	}

	err = c.CurrencyRepository.UpInsert(ctx, entities.Currency{Code: resp.Code, Name: name, USDConversionRate: rate})
	return err
}

func shouldUpdateCurrencyData(currencyData entities.Currency, err error) bool {
	return err == mongo.ErrNoDocuments || time.Now().After(currencyData.UpdatedAt.Add(30*time.Second))
}
