package users_postgres_repositpory

import "github.com/Aam-Shaegar/todo-list/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))
	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.PhoneNumber,
			user.FullName,
		)
	}
	return userDomains
}
