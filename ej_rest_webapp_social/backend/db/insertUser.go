package db

import (
	"github.com/zeroidentidad/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertUser para insertar los datos del usuario en DB*/
func InsertUser(user models.User) (string, bool, error) {
	ctx, cancel := Timeout()
	defer cancel()

	db := MongoConn.Database("twittor")
	col := db.Collection("usuarios")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
