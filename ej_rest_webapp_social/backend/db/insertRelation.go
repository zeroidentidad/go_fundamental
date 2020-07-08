package db

import "github.com/zeroidentidad/twittor/models"

/*InsertRelation graba la relación en la BD */
func InsertRelation(relation models.Relation) (bool, error) {
	ctx, cancel := Timeout()
	defer cancel()

	db := MongoConn.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, relation)
	if err != nil {
		return false, err
	}

	return true, nil
}
