package application

import (
	"github.com/srhtak/go-ddd-api/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByUsername(username string) (*domain.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(username, password string) error {
	user, err := domain.NewUser(username, password)
	if err != nil {
		return err
	}
	return s.repo.Create(user)
}

func (s *UserService) AuthenticateUser(username, password string) bool {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return false
	}
	return user.ValidatePassword(password)
}