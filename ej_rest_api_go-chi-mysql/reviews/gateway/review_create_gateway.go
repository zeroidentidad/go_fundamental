package gateway

import (
	"github.com/zeroidentidad/rest-chi-mysql/internal/database"
	"github.com/zeroidentidad/rest-chi-mysql/reviews/models"
)

type ReviewGateway interface {
	AddReview(cmd *models.CreateReviewCMD) (string, error)
}

type ReviewGtw struct {
	ReviewStorage
}

func NewReviewGateway(mongoClient *database.Mongo) ReviewGateway {
	return &ReviewGtw{&ReviewStg{mongoClient}}
}
