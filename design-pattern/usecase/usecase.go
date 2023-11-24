package usecase

import (
	"context"

	"github.com/kurakura967/go-layered-architecture-anti-pattern/design-pattern/infra"
)

type UserDTO struct {
	Id   int
	Name string
}

type UserRepositorier interface {
	Get(ctx context.Context, userId int) (infra.User, error)
}

func GetUserById(ctx context.Context, userId int, repo UserRepositorier) (UserDTO, error) {

	user, err := repo.Get(ctx, userId)
	if err != nil {
		// error handling
	}

	return UserDTO{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}
