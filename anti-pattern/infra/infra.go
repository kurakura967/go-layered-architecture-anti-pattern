package infra

import (
	"context"
	"database/sql"
)

type User struct {
	Id   int
	Name string
}

func Get(ctx context.Context, userId int, db *sql.DB) (User, error) {
	u := User{}

	err := db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", userId).Scan(&u.Id, &u.Name)
	if err != nil {
		// error handling
	}
	return u, nil
}
