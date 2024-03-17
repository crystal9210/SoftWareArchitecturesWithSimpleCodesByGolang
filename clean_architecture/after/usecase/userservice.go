package usecase

import "architectures/clean_architecture/after/entity"

type UserRepository interface {
	GetAllUsers() []entity.User
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() []entity.User {
	return s.repo.GetAllUsers()
}
