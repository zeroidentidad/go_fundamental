package db

import (
	"../models"
)

func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.001,
	}
}
