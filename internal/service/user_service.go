package service

import (
	"context"

	"project/internal/model"
	"project/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	return s.repo.Create(ctx, user)
}

func (s *UserService) GetUsers(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAll(ctx)
}