package users_postgres_repositpory

import (
	"context"
	"errors"
	"fmt"

	"github.com/Aam-Shaegar/todo-list/internal/core/domain"
	core_error "github.com/Aam-Shaegar/todo-list/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) GetUser(ctx context.Context, userID int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, full_name, phone_number
		FROM todoapp.users
		WHERE id=$1
	`
	row := r.pool.QueryRow(ctx, query, userID)
	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id='%d' : %w", userID, core_error.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)

	}
	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.PhoneNumber,
		userModel.FullName,
	)
	return userDomain, nil
}
