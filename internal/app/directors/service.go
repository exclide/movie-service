package directors

import (
	"context"
	"github.com/exclide/movie-service/internal/app/model"
)

type Service interface {
	Create(ctx context.Context, e *model.Director) (*model.Director, error)
	GetById(ctx context.Context, id int) (*model.Director, error)
	DeleteById(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.Director, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(ctx context.Context, e *model.Director) (*model.Director, error) {
	m, err := s.repo.Create(ctx, e)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *service) GetById(ctx context.Context, id int) (*model.Director, error) {
	m, err := s.repo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *service) DeleteById(ctx context.Context, id int) error {
	err := s.repo.DeleteById(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAll(ctx context.Context) ([]model.Director, error) {
	m, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return m, nil
}
