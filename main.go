package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/external"
	"github.com/lumacielz/challenge-bravo/handlers"
	"github.com/lumacielz/challenge-bravo/presenters"
	"github.com/lumacielz/challenge-bravo/useCases"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

func init() {
	viper.SetConfigFile("config.json")
	viper.ReadInConfig()
}

func main() {
	const baseUrl = "https://economia.awesomeapi.com.br/json/%s-USD"
	databaseCfg := viper.GetStringMapString("database")
	opts := options.Client().ApplyURI(databaseCfg["uri"])
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database(databaseCfg["name"]).Collection(databaseCfg["collection"])

	mongoClient := database.Client{Collection: collection}
	quotationAPICLient := external.QuotationClient{
		Url:     viper.GetString("external.quotationAPI.url"),
		Timeout: viper.GetDuration("external.quotationAPI.timeout"),
		Client:  http.DefaultClient,
	}

	currencyUseCase := useCases.CurrencyUseCase{
		Now:                 func() time.Time { return time.Now() },
		UpdateFrequency:     viper.GetDuration("updateFrequency"),
		CurrencyRepository:  mongoClient,
		QuotationRepository: quotationAPICLient,
	}

	currencyController := handlers.CurrencyController{
		UseCase:         currencyUseCase,
		OutputPresenter: presenters.JsonPresenter{},
		InputPresenter:  presenters.JsonPresenter{},
		Timeout:         viper.GetDuration("timeout"),
	}

	r := chi.NewRouter()

	r.Get("/currency", currencyController.GetConversionHandler)
	r.Post("/currency/new", currencyController.NewCurrencyHandler)
	r.Put("/currency/{code}", currencyController.UpdateCurrencyHandler)
	r.Delete("/currency/{code}", currencyController.DeleteCurrencyHandler)

	http.ListenAndServe(":8080", r)
}
