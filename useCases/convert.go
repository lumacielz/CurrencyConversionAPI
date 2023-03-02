package useCases

import (
	"context"
	"fmt"
	"github.com/lumacielz/challenge-bravo/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
	"strconv"
	"strings"
	"time"
)

func (c CurrencyUseCase) Convert(ctx context.Context, amount float64, from, to string) (CurrencyConversionResponse, error) {
	originCurrencyDataC := make(chan entities.Currency, 1)
	destinationCurrencyDataC := make(chan entities.Currency, 1)

	defer close(originCurrencyDataC)
	defer close(destinationCurrencyDataC)

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		originCurrencyData, err := c.CurrencyRepository.Get(gCtx, from)
		if c.shouldUpdateCurrencyData(originCurrencyData, err) {
			c.UpdateCurrencyData(ctx, from)
			originCurrencyData, err = c.CurrencyRepository.Get(gCtx, from)
		}

		originCurrencyDataC <- originCurrencyData
		return err
	})
	g.Go(func() error {
		destinationCurrencyData, err := c.CurrencyRepository.Get(ctx, to)
		if c.shouldUpdateCurrencyData(destinationCurrencyData, err) {
			c.UpdateCurrencyData(ctx, to)
			destinationCurrencyData, err = c.CurrencyRepository.Get(ctx, to)
		}

		destinationCurrencyDataC <- destinationCurrencyData
		return err
	})

	err := g.Wait()
	if err != nil {
		return CurrencyConversionResponse{}, err
	}

	originCurrencyData := <-originCurrencyDataC
	destinationCurrencyData := <-destinationCurrencyDataC

	destinationValue := convert(originCurrencyData.USDConversionRate, destinationCurrencyData.USDConversionRate, amount)
	response := CurrencyConversionResponse{
		Value:    fmt.Sprintf("%.3f", destinationValue),
		Currency: to,
	}

	return response, nil
}

func (c CurrencyUseCase) shouldUpdateCurrencyData(currencyData entities.Currency, err error) bool {
	return err == mongo.ErrNoDocuments || c.Now().After(currencyData.UpdatedAt.Add(30*time.Second))
}

func (c CurrencyUseCase) UpdateCurrencyData(ctx context.Context, code string) error {
	resp, err := c.QuotationRepository.GetCurrentUSDQuotation(ctx, code)

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

func convert(originRate, destinationRate, amount float64) float64 {
	return originRate / destinationRate * amount
}
