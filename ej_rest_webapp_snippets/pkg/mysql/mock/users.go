package mock

import (
	"pastein/pkg/models"
	"time"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "Test",
	Email:   "test@example.com",
	Created: time.Now(),
}

type UserModel struct{}

func (m *UserModel) Insert(r *models.UserRequest) error {
	switch r.Email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(r *models.UserRequest) (int, error) {
	switch r.Email {
	case "test@example.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
