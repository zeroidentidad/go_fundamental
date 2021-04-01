package mock

import (
	"pastein/pkg/models"
	"time"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "Un viejo estanque silencioso",
	Content: "Un viejo estanque silencioso...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(r *models.SnippetRequest) (int, error) {
	return 2, nil

}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
