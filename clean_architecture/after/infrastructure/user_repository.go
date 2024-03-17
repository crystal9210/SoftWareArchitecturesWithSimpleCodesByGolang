package infrastructure

import "architectures/clean_architecture/after/entity"

type InMemoryUserRepository struct {
	users []entity.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	// Initialize with some users
	users := []entity.User{{ID: 1, Name: "John Doe"}}
	return &InMemoryUserRepository{users: users}
}

func (repo *InMemoryUserRepository) GetAllUsers() []entity.User {
	return repo.users
}
