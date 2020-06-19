package db

import (
	"context"
	"fmt"
	"log"

	"github.com/zeroidentidad/twittor/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConn = connectDB()

var confObj, connString = *config.GetConfiguration(), fmt.Sprintf("%s", confObj.DBUrl)

var clientOptions = options.Client().ApplyURI(connString)

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("DB connection success...")

	return client
}

func CheckConnection() int {
	err := MongoConn.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
