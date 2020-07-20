package db

import (
	"github.com/zeroidentidad/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadAllUsers leer los usuarios registrados en el sistema, si se recibe "R" en quienes trae solo los que se relacionan conmigo */
func ReadAllUsers(ID string, page int64, search string, tipo string) ([]*models.User, bool) {
	ctx, cancel := Timeout()
	defer cancel()

	db := MongoConn.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return results, false
	}

	var encontrado, incluir bool

	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return results, false
		}

		var relation models.Relation
		relation.UserID = ID
		relation.UserIDRelation = user.ID.Hex()

		incluir = false

		encontrado, err = SearchRelation(relation)
		if tipo == "new" && encontrado == false {
			incluir = true
		}
		if tipo == "follow" && encontrado == true {
			incluir = true
		}

		if relation.UserIDRelation == ID {
			incluir = false
		}

		if incluir == true {
			user.Password = ""
			user.Biografia = ""
			user.SitioWeb = ""
			user.Ubicacion = ""
			user.Banner = ""
			user.Email = ""

			results = append(results, &user)
		}
	}

	err = cursor.Err()
	if err != nil {
		return results, false
	}
	cursor.Close(ctx)
	return results, true
}
