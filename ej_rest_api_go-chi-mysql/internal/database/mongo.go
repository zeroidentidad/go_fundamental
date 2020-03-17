package database

import (
	"context"
	"fmt"
	"time"

	"github.com/zeroidentidad/rest-chi-mysql/internal/logs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	Client *mongo.Client
}

func NewMongoClient(host string) *Mongo {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", host)))

	if err != nil {
		logs.Log().Errorf("Could not connect with mongo %s", err.Error())
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		logs.Log().Errorf("Could not connect with mongo %s", err.Error())
		panic(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		logs.Log().Errorf("Could not PING the mongo server: %s", err.Error())
		panic(err)
	}

	return &Mongo{client}
}
