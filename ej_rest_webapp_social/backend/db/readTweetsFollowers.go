package db

import (
	"github.com/zeroidentidad/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ReadTweetsFollowers lee los tweets de seguidores*/
func ReadTweetsFollowers(ID string, pagina int) ([]models.ReturnTweetsFollowers, bool) {
	ctx, cancel := Timeout()
	defer cancel()

	db := MongoConn.Database("twittor")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"userid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "useridrelation",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)

	var result []models.ReturnTweetsFollowers
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
