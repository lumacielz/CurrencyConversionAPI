package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/lumacielz/challenge-bravo/database"
	"github.com/lumacielz/challenge-bravo/external"
	"github.com/lumacielz/challenge-bravo/handlers"
	"github.com/lumacielz/challenge-bravo/presenters"
	"github.com/lumacielz/challenge-bravo/useCases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func main() {
	const uri = "mongodb://challenge-bravo:bravo123@localhost:27017"

	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("challenge-bravo").Collection("currencies")

	mongoClient := database.Client{Collection: collection}
	currencyUseCase := useCases.CurrencyUseCase{CurrencyRepository: mongoClient, QuotationRepository: external.QuotationClient{http.DefaultClient}}
	currencyController := handlers.CurrencyController{UseCase: currencyUseCase, OutputPresenter: presenters.JsonPresenter{}, InputPresenter: presenters.JsonPresenter{}}

	r := chi.NewRouter()

	r.Get("/currency", currencyController.GetConversionHandler)
	r.Post("/currency/new", currencyController.NewCurrencyHandler)
	r.Put("/currency/{code}", currencyController.UpdateCurrencyHandler)
	r.Delete("/currency/{code}", currencyController.DeleteCurrencyHandler)

	http.ListenAndServe(":8080", r)
}
