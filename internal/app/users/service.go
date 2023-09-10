package users

import (
	"context"
	"github.com/exclide/movie-service/internal/app/model"
)

type Service interface {
	Create(ctx context.Context, e *model.User) (*model.User, error)
	GetByLogin(ctx context.Context, login string) (*model.User, error)
	DeleteByLogin(ctx context.Context, login string) error
	GetAll(ctx context.Context) ([]model.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(ctx context.Context, e *model.User) (*model.User, error) {
	m, err := s.repo.Create(ctx, e)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *service) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	m, err := s.repo.GetByLogin(ctx, login)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *service) DeleteByLogin(ctx context.Context, login string) error {
	err := s.repo.DeleteByLogin(ctx, login)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAll(ctx context.Context) ([]model.User, error) {
	m, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return m, nil
}
