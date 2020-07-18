package db

import (
	"github.com/zeroidentidad/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*SearchRelation consultar relacion entre 2 usuarios*/
func SearchRelation(relation models.Relation) (bool, error) {
	ctx, cancel := Timeout()
	defer cancel()

	db := MongoConn.Database("twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"userid":         relation.UserID,
		"useridrelation": relation.UserIDRelation,
	}

	var resultado models.Relation
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, err
	}
	return true, nil
}
