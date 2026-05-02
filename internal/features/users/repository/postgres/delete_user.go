package users_postgres_repositpory

import (
	"context"
	"fmt"

	core_error "github.com/Aam-Shaegar/todo-list/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE FROM todoapp.users
		WHERE id=$1;
	`
	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("users with id='%d': %w", id, core_error.ErrNotFound)
	}
	return nil
}
