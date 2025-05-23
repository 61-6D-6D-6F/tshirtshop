package service

import (
	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
	"github.com/61-6D-6D-6F/tshirtshop/internal/repository"
)

type TShirtService interface {
	ListTShirts() ([]*model.TShirt, error)
	GetTShirt(int) (*model.TShirt, error)
	CreateTShirt(*model.TShirt) error
	UpdateTShirt(int, *model.TShirt) error
	DeleteTShirt(int) error
}

type tShirtService struct {
	repo repository.TShirtRepository
}

func NewTShirtService(r repository.TShirtRepository) TShirtService {
	return &tShirtService{repo: r}
}

func (s *tShirtService) ListTShirts() ([]*model.TShirt, error) {
	return s.repo.List()
}

func (s *tShirtService) GetTShirt(id int) (*model.TShirt, error) {
	return s.repo.Get(id)
}

func (s *tShirtService) CreateTShirt(tShirt *model.TShirt) error {
	return s.repo.Save(tShirt)
}

func (s *tShirtService) UpdateTShirt(id int, tShirt *model.TShirt) error {
	return s.repo.Update(id, tShirt)
}

func (s *tShirtService) DeleteTShirt(id int) error {
	return s.repo.Delete(id)
}
