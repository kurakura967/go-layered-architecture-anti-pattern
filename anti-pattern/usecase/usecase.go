package usecase

import (
	"context"
	"database/sql"

	"github.com/kurakura967/go-layered-architecture-anti-pattern/anti-pattern/infra"
)

type UserDTO struct {
	Id   int
	Name string
}

func GetUserById(ctx context.Context, userId int, db *sql.DB) (UserDTO, error) {
	//db, err := connectDB()
	//if err != nil {
	//	// error handling
	//}
	//defer db.Close()

	user, err := infra.Get(ctx, userId, db)
	if err != nil {
		// error handling
	}
	return UserDTO{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/sample")
	if err != nil {
		return nil, err
	}
	return db, nil
}
