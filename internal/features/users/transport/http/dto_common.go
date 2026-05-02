package users_transport_http

import "github.com/Aam-Shaegar/todo-list/internal/core/domain"

type UserResponseDto struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userDtoFromDomain(user domain.User) UserResponseDto {
	return UserResponseDto{
		ID:          user.ID,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		Version:     user.Version,
	}
}

func usersDtoFromDomains(users []domain.User) []UserResponseDto {
	usersDto := make([]UserResponseDto, len(users))
	for i, user := range users {
		usersDto[i] = userDtoFromDomain(user)
	}
	return usersDto
}
