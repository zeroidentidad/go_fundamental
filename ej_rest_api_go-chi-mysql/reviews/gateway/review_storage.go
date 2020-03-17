package gateway

import (
	"context"
	"time"

	"github.com/zeroidentidad/rest-chi-mysql/internal/database"
	"github.com/zeroidentidad/rest-chi-mysql/internal/logs"
	"github.com/zeroidentidad/rest-chi-mysql/reviews/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewStorage interface {
	Add(cmd *models.CreateReviewCMD) (string, error)
}

type ReviewStg struct {
	*database.Mongo
}

func (s *ReviewStg) Add(cmd *models.CreateReviewCMD) (string, error) {
	coll := s.Client.Database("reviewDB").Collection("reviews")

	res, err := coll.InsertOne(context.Background(),
		bson.D{
			{"comment", cmd.Comment},
			{"stars", cmd.Stars},
			{"createdAt", time.Now()},
			{"gadgetId", cmd.GadgetId},
		})

	if err != nil {
		logs.Log().Error("cannot insert review")
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID)

	return id.String(), nil
}
