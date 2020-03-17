package gateway

import "github.com/zeroidentidad/rest-chi-mysql/gadgets/smartphones/models"

func (s *SmartphoneCreateGtw) Create(cmd *models.CreateSmartphoneCMD) (*models.Smartphone, error) {
	return s.create(cmd)
}
