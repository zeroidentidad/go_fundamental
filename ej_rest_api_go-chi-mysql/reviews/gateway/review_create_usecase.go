package gateway

import "github.com/zeroidentidad/rest-chi-mysql/reviews/models"

func (g *ReviewGtw) AddReview(cmd *models.CreateReviewCMD) (string, error) {
	return g.Add(cmd)
}
