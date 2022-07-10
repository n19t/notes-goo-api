package db

import (
	"context"
	"fmt"
	"notes-goo-api/internal/config"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(c config.ConfigDb) (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s%s",
		strings.Replace(c.Username, "@", "%40", -1),
		strings.Replace(c.Password, "@", "%40", -1),
		c.Host,
		c.Database)
	// "mongodb+srv://admin:P%40ssw0rd@cluster0.0z4ce.mongodb.net"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		uri,
	))
	if err != nil {
		return nil, err
	}
	return client, nil
}
