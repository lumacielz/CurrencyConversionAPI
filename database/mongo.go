package database

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Client struct {
	Collection *mongo.Collection
}

func (c Client) Get(ctx context.Context, code string) (entities.Currency, error) {
	var currency entities.Currency
	err := c.Collection.FindOne(ctx, bson.M{"code": code}).Decode(&currency)
	if err == mongo.ErrNoDocuments {
		err = entities.ErrCurrencyNotFound
	}
	return currency, err
}

func (c Client) Create(ctx context.Context, currency entities.Currency) (interface{}, error) {
	currency.UpdatedAt = time.Now()
	res, err := c.Collection.InsertOne(ctx, currency)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func (c Client) UpInsert(ctx context.Context, currency entities.Currency) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"code": currency.Code}

	currency.UpdatedAt = time.Now()

	_, err := c.Collection.UpdateOne(ctx, filter, bson.M{"$set": currency}, opts)
	return err
}

func (c Client) Update(ctx context.Context, code string, currency entities.Currency) error {
	filter := bson.M{"code": code}

	updatedCurrency := bson.M{}
	if currency.Name != "" {
		updatedCurrency["name"] = currency.Name
	}
	updatedCurrency["USDConvertionRate"] = currency.USDConversionRate
	updatedCurrency["updatedAt"] = time.Now()

	res, err := c.Collection.UpdateOne(ctx, filter, bson.M{"$set": updatedCurrency})
	if res.MatchedCount == 0 {
		return entities.ErrCurrencyNotFound
	}

	return err
}

func (c Client) Delete(ctx context.Context, code string) error {
	filter := bson.M{"code": code}
	res, err := c.Collection.DeleteOne(ctx, filter)

	if res.DeletedCount == 0 {
		return entities.ErrCurrencyNotFound
	}

	return err
}
