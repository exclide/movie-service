package directors

import (
	"context"
	"github.com/exclide/movie-service/internal/app/model"
)

type Repository interface {
	Create(ctx context.Context, director *model.Director) (*model.Director, error)
	GetById(ctx context.Context, id int) (*model.Director, error)
	DeleteById(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.Director, error)
}
