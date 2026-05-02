package users_service

import (
	"context"
	"fmt"

	"github.com/Aam-Shaegar/todo-list/internal/core/domain"
)

func (s *UsersService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}
	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("repository CreateUser error: %w", err)
	}
	return user, nil
}
