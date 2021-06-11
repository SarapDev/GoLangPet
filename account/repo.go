package account

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handl request")

type repo struct {
	db		*sql.DB
	logger 	log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db: db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (r repo) CreateUser(context context.Context, user User) error {
	sqlQuery := `
		INSERT INTO users (id, email, password)
		VALUES (?, ?, ?)`

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	_, err := r.db.ExecContext(context, sqlQuery, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r repo) GetUser(context context.Context, id string) (string, error) {
	var email string

	err := r.db.QueryRow(`SELECT email FROM users WHERE id=?`, id).Scan(&email)
	if err != nil {
		return "", RepoErr
	}

	return email, nil
}