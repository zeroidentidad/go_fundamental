package db

import (
	"context"

	"github.com/zeroidentidad/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadTweets lee los tweets de un perfil */
func ReadTweets(ID string, pagina int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := Timeout()
	defer cancel()

	db := MongoConn.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.ReturnTweets

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultados, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.ReturnTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
