package repository

import (
	"context"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/internal/app/store"
)

type DirectorRepository struct {
	Store *store.Store
}

func NewDirectorRepository(s *store.Store) DirectorRepository {
	return DirectorRepository{s}
}

func (r *DirectorRepository) Create(ctx context.Context, u *model.Director) (*model.Director, error) {
	stmt, err := r.Store.Db.Prepare("INSERT INTO directors (name) VALUES ($1) RETURNING id")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(u.Name).Scan(&u.Id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *DirectorRepository) GetById(ctx context.Context, id int) (*model.Director, error) {
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

func (r *DirectorRepository) DeleteById(ctx context.Context, id int) error {
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

func (r *DirectorRepository) GetAll(ctx context.Context) ([]model.Director, error) {
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
