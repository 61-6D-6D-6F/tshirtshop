package service

import (
	"github.com/61-6D-6D-6F/tshirtshop/internal/model"
	"github.com/61-6D-6D-6F/tshirtshop/internal/repository"
)

type UserService interface {
	ListUsers() ([]*model.User, error)
	GetUser(int) (*model.User, error)
	CreateUser(*model.User) error
	UpdateUser(int, *model.User) error
	DeleteUser(int) error
	TryLogin(string, string) (*model.User, error)
	TryRegister(*model.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) ListUsers() ([]*model.User, error) {
	return s.repo.List()
}

func (s *userService) GetUser(id int) (*model.User, error) {
	return s.repo.Get(id)
}

func (s *userService) CreateUser(tShirt *model.User) error {
	return s.repo.Save(tShirt)
}

func (s *userService) UpdateUser(id int, user *model.User) error {
	return s.repo.Update(id, user)
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}

func (s *userService) TryLogin(username string, password string) (*model.User, error) {
	return s.repo.TryLogin(username, password)
}

func (s *userService) TryRegister(user *model.User) error {
	return s.repo.TryRegister(user)
}
