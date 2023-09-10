package movies

import (
	"context"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/internal/app/store"
)

type Repository interface {
	Create(ctx context.Context, e *model.Movie) (*model.Movie, error)
	GetById(ctx context.Context, id int) (*model.Movie, error)
	DeleteById(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.Movie, error)
}

type repository struct {
	Store *store.Store
}

func NewRepository(s *store.Store) Repository {
	return &repository{s}
}

func (r *repository) Create(ctx context.Context, e *model.Movie) (*model.Movie, error) {
	stmt, err := r.Store.Db.Prepare("INSERT INTO movies (title, year, director_id) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(e.Title, e.Year, e.DirectorId).Scan(&e.Id)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *repository) GetById(ctx context.Context, id int) (*model.Movie, error) {
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

func (r *repository) DeleteById(ctx context.Context, id int) error {
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

func (r *repository) GetAll(ctx context.Context) ([]model.Movie, error) {
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
