package infra

import (
	"context"
	"database/sql"
)

type User struct {
	Id   int
	Name string
}

type sqlHandler struct {
	db *sql.DB
}

func NewSqlHandler(db *sql.DB) *sqlHandler {
	return &sqlHandler{db: db}
}

func (s *sqlHandler) Get(ctx context.Context, userId int) (User, error) {
	u := User{}

	err := s.db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", userId).Scan(&u.Id, &u.Name)
	if err != nil {
		// error handling
	}
	return u, nil
}
