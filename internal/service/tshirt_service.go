package service

import (
	"errors"

	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
	"github.com/61-6D-6D-6F/tshirtshop/internal/repository"
)

type TShirtService interface {
	ListTShirts() ([]model.TShirt, error)
	CreateTShirt(model.TShirt) error
	UpdateTShirt(int, model.TShirt) error
	DeleteTShirt(int) error
}

type tShirtService struct {
	repo repository.TShirtRepository
}

func NewTShirtService(r repository.TShirtRepository) TShirtService {
	return &tShirtService{repo: r}
}

func (s *tShirtService) ListTShirts() ([]model.TShirt, error) {
	return s.repo.List()
}

func (s *tShirtService) CreateTShirt(tShirt model.TShirt) error {
	if tShirt.Name == "" || tShirt.Size == "" || tShirt.Color == "" ||
		tShirt.Price == 0.0 || tShirt.Stock == 0 {
		return errors.New("error: Invalid input")
	}
	return s.repo.Save(tShirt)
}

func (s *tShirtService) UpdateTShirt(id int, tShirt model.TShirt) error {
	if tShirt.Name == "" || tShirt.Size == "" || tShirt.Color == "" ||
		tShirt.Price == 0.0 || tShirt.Stock == 0 {
		return errors.New("error: Invalid input")
	}
	return s.repo.Update(id, tShirt)
}

func (s *tShirtService) DeleteTShirt(id int) error {
	return s.repo.Delete(id)
}
