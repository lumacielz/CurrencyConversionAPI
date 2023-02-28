package useCases

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
	"strconv"
	"strings"
	"time"
)

func (c *CurrencyUseCase) Convert(ctx context.Context, amount float64, from, to string) (CurrencyResponse, error) {
	originCurrencyDataC := make(chan entities.Currency, 1)
	destinationCurrencyDataC := make(chan entities.Currency, 1)

	defer close(originCurrencyDataC)
	defer close(destinationCurrencyDataC)

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		originCurrencyData, err := c.CurrencyRepository.Get(gCtx, from)
		if shouldUpdateCurrencyData(originCurrencyData, err) {
			c.UpdateCurrencyData(ctx, from)
			originCurrencyData, err = c.CurrencyRepository.Get(gCtx, from)
		}

		originCurrencyDataC <- originCurrencyData
		return err
	})
	g.Go(func() error {
		destinationCurrencyData, err := c.CurrencyRepository.Get(ctx, to)
		if shouldUpdateCurrencyData(destinationCurrencyData, err) {
			c.UpdateCurrencyData(ctx, to)
			destinationCurrencyData, err = c.CurrencyRepository.Get(ctx, to)
		}

		destinationCurrencyDataC <- destinationCurrencyData
		return err
	})

	err := g.Wait()
	if err != nil {
		return CurrencyResponse{}, err
	}

	originCurrencyData := <-originCurrencyDataC
	destinationCurrencyData := <-destinationCurrencyDataC

	destinationValue := convert(originCurrencyData.USDConversionRate, destinationCurrencyData.USDConversionRate, amount)
	response := CurrencyResponse{
		Value:    destinationValue,
		Currency: to,
	}

	return response, nil
}

//TODO revisar se faz sentido ignorar erro
func (c *CurrencyUseCase) UpdateCurrencyData(ctx context.Context, code string) error {
	resp, err := c.QuotationClient.GetCurrentUSDQuotation(ctx, code)

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

func convert(originRate, destinationRate, amount float64) float64 {
	return originRate / destinationRate * amount
}
