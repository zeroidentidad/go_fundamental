package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ReturnTweetsFollowers es la estructura con la que devolveremos los tweets */
type ReturnTweetsFollowers struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userid,omitempty"`
	UserIDRelation string             `bson:"useridrelation" json:"useridrelation,omitempty"`
	Tweet          struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
