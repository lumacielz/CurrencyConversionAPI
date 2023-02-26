package database

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	Database *mongo.Collection
}

//TODO implementar repository em Client
func (c Client) Get(ctx context.Context, code string) (entities.Currency, error) {
	var currency entities.Currency
	err := c.Database.FindOne(ctx, bson.D{{"code", code}}).Decode(&currency)
	return currency, err
}
