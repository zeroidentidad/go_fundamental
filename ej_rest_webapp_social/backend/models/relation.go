package models

/* Relation data model para guardar relaion de un usuario con otro*/
type Relation struct {
	UserID         string `bson:"userid" json:"userid"`
	UserIDRelation string `bson:"useridrelation" json:"useridrelation"`
}
