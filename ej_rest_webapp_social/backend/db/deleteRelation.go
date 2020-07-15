package db

import "github.com/zeroidentidad/twittor/models"

/*DeleteRelation borrar relacion usurariosId en DB*/
func DeleteRelation(relation models.Relation) (bool, error) {
	ctx, cancel := Timeout()
	defer cancel()

	db := MongoConn.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, relation)
	if err != nil {
		return false, err
	}
	return true, nil
}
