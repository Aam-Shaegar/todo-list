package users_service

import (
	"context"

	"github.com/Aam-Shaegar/todo-list/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, userID int) (domain.User, error)
	DeleteUser(ctx context.Context, userID int) error
	PatchUser(ctx context.Context, id int, user domain.User) (domain.User, error)
}

func NewUsersService(usersRepository UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}
