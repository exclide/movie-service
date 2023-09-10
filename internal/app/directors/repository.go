package directors

import (
	"context"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/internal/app/store"
)

type Repository interface {
	Create(ctx context.Context, e *model.Director) (*model.Director, error)
	GetById(ctx context.Context, id int) (*model.Director, error)
	DeleteById(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.Director, error)
}

type repository struct {
	Store *store.Store
}

func NewRepository(s *store.Store) Repository {
	return &repository{s}
}

func (r *repository) Create(ctx context.Context, e *model.Director) (*model.Director, error) {
	stmt, err := r.Store.Db.Prepare("INSERT INTO directors (name) VALUES ($1) RETURNING id")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(e.Name).Scan(&e.Id)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *repository) GetById(ctx context.Context, id int) (*model.Director, error) {
	var mv model.Director

	stmt, err := r.Store.Db.Prepare("select * from directors where id = $1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&mv.Id, &mv.Name)
	if err != nil {
		return nil, err
	}

	return &mv, nil
}

func (r *repository) DeleteById(ctx context.Context, id int) error {
	stmt, err := r.Store.Db.Prepare("delete from directors where id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAll(ctx context.Context) ([]model.Director, error) {
	var mvs []model.Director

	stmt, err := r.Store.Db.Prepare("select * from directors")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	defer rows.Close()

	for rows.Next() {
		var mv model.Director
		if err := rows.Scan(&mv.Id, &mv.Name); err != nil {
			return nil, err
		}
		mvs = append(mvs, mv)
	}

	return mvs, nil
}
