package service

import "cart/internal/models"

type CartService struct {
}

func NewCartService() *CartService {
	return &CartService{}
}

func (c *CartService) Create(userid string, currencies map[string]int) (string, error) {
	return "created", nil
}

func (c *CartService) Get(id string) (models.Cart, error) {
	return models.Cart{
		UserId:   id,
		Products: map[string]int{},
	}, nil
}
