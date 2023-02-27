package database

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	Collection *mongo.Collection
}

func (c Client) Get(ctx context.Context, code string) (entities.Currency, error) {
	var currency entities.Currency
	err := c.Collection.FindOne(ctx, bson.D{{"code", code}}).Decode(&currency)
	return currency, err
}
