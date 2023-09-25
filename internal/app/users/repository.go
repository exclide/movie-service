package users

import (
	"context"
	"database/sql"
	"github.com/exclide/movie-service/internal/app/model"
	"github.com/exclide/movie-service/internal/app/store"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Create(ctx context.Context, e *model.User) (*model.User, error)
	GetByLogin(ctx context.Context, login string) (*model.User, error)
	DeleteByLogin(ctx context.Context, login string) error
	GetAll(ctx context.Context) ([]model.User, error)
}

type repository struct {
	Store *store.Store
}

func NewRepository(s *store.Store) Repository {
	return &repository{s}
}

func (r *repository) Create(ctx context.Context, e *model.User) (*model.User, error) {
	stmt, err := r.Store.Db.Prepare("INSERT INTO users (login, password) VALUES ($1, $2)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(e.Login, e.Password)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *repository) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	var mv model.User

	stmt, err := r.Store.Db.Prepare("select * from users where login = $1")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(login).Scan(&mv.Login, &mv.Password)
	if err != nil {
		return nil, err
	}

	return &mv, nil
}

func (r *repository) DeleteByLogin(ctx context.Context, login string) error {
	stmt, err := r.Store.Db.Prepare("delete from users where login = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(login)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAll(ctx context.Context) ([]model.User, error) {
	var mvs []model.User

	stmt, err := r.Store.Db.Prepare("select * from users")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			logrus.Info(err)
		}
	}(rows)

	for rows.Next() {
		var mv model.User
		if err := rows.Scan(&mv.Login, &mv.Password); err != nil {
			return nil, err
		}
		mvs = append(mvs, mv)
	}

	return mvs, nil
}
