package database

import (
	"context"
	"github.com/lumacielz/challenge-bravo/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Client struct {
	Collection *mongo.Collection
}

func (c *Client) Get(ctx context.Context, code string) (entities.Currency, error) {
	var currency entities.Currency
	err := c.Collection.FindOne(ctx, bson.D{{"code", code}}).Decode(&currency)
	return currency, err
}

func (c *Client) Create(ctx context.Context, currency entities.Currency) error {
	currency.UpdatedAt = time.Now()
	_, err := c.Collection.InsertOne(ctx, currency)
	return err
}

func (c *Client) Update(ctx context.Context, currency entities.Currency) error {
	currency.UpdatedAt = time.Now()
	filter := bson.D{{"code", currency.Code}}
	_, err := c.Collection.UpdateOne(ctx, filter, currency)
	return err
}

func (c *Client) Delete(ctx context.Context, code string) error {
	filter := bson.D{{"code", code}}
	_, err := c.Collection.DeleteOne(ctx, filter)
	return err
}
