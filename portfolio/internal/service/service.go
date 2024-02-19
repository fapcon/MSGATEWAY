package service

import "portfolio/internal/models"

type PortfService struct {
}

func NewPortfService() *PortfService {
	return &PortfService{}
}

func (p *PortfService) Create(userid string, currencies map[string]float64) (string, error) {
	return "created", nil
}

func (p *PortfService) Get(id string) (models.Portfolio, error) {
	return models.Portfolio{
		UserId:     id,
		Currencies: map[string]float64{},
	}, nil
}
