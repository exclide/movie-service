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
	stmt, err := r.Store.Db.Prepare("INSERT INTO movies (title, year, director_id) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(u.Title, u.Year, u.DirectorId).Scan(&u.Id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *MovieRepository) GetById(ctx context.Context, id int) (*model.Movie, error) {
	var mv model.Movie

	stmt, err := r.Store.Db.Prepare("select * from movies where id = $1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&mv.Id, &mv.Title, &mv.Year, &mv.DirectorId)
	if err != nil {
		return nil, err
	}

	return &mv, nil
}

func (r *MovieRepository) DeleteById(ctx context.Context, id int) error {
	stmt, err := r.Store.Db.Prepare("delete from movies where id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *MovieRepository) GetAll(ctx context.Context) ([]model.Movie, error) {
	var mvs []model.Movie

	stmt, err := r.Store.Db.Prepare("select * from movies")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	defer rows.Close()

	for rows.Next() {
		var mv model.Movie
		if err := rows.Scan(&mv.Id, &mv.Title,
			&mv.Year, &mv.DirectorId); err != nil {
			return nil, err
		}
		mvs = append(mvs, mv)
	}

	return mvs, nil
}
