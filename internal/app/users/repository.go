package users

import (
	"context"
	"github.com/exclide/movie-service/internal/app/model"
)

type Repository interface {
	Create(ctx context.Context, movie *model.User) (*model.User, error)
	GetByLogin(ctx context.Context, login string) (*model.User, error)
	DeleteByLogin(ctx context.Context, login string) error
	GetAll(ctx context.Context) ([]model.User, error)
}
