package movies

import (
	"context"
	"github.com/exclide/movie-service/internal/app/model"
)

type Repository interface {
	Create(ctx context.Context, movie *model.Movie) (*model.Movie, error)
	GetById(ctx context.Context, id int) (*model.Movie, error)
	DeleteById(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.Movie, error)
}
