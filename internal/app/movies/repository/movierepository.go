package repository

import (
	"context"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/internal/app/store"
)

type MovieRepository struct {
	Store *store.Store
}

func NewMovieRepository(s *store.Store) MovieRepository {
	return MovieRepository{s}
}

func (r *MovieRepository) Create(ctx context.Context, u *model.Movie) (*model.Movie, error) {
	if err := r.Store.Db.QueryRow(
		"INSERT INTO movies (title, year, director_id) VALUES ($1, $2, $3) RETURNING id",
		u.Title,
		u.Year,
		u.DirectorId,
	).Scan(&u.Id); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *MovieRepository) GetById(ctx context.Context, id int) (*model.Movie, error) {
	var mv model.Movie

	err := r.Store.Db.QueryRow("select * from movies where id = $1", id).
		Scan(&mv.Id, &mv.Title, &mv.Year, &mv.DirectorId)
	if err != nil {
		return nil, err
	}

	return &mv, nil
}

func (r *MovieRepository) DeleteById(ctx context.Context, id int) error {
	return nil
}

func (r *MovieRepository) GetAll(ctx context.Context) ([]model.Movie, error) {
	return nil, nil
}
