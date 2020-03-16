package models

import "testing"

func NewReview(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   stars,
		Comment: comment,
	}
}

func Test_withCorrectParams(t *testing.T) {
	r := NewReview(4, "El iphone X es una mierda")

	err := r.validate()

	if err != nil {
		t.Error("the validation did not pass")
		t.Fail()
	}
}

func Test_shouldFailWithWrongNumberOfStars(t *testing.T) {
	r := NewReview(8, "El iphone parece ser bueno")

	err := r.validate()

	if err == nil {
		t.Error("should fail with 5 stars")
		t.Fail()
	}
}
