package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DeleteTweet borra un tweet determinado */
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := Timeout()
	defer cancel()

	db := MongoConn.Database("twittor")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}
