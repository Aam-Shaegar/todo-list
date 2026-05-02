package users_transport_http

import (
	"net/http"

	"github.com/Aam-Shaegar/todo-list/internal/core/domain"
	core_logger "github.com/Aam-Shaegar/todo-list/internal/core/logger"
	core_http_request "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/request"
	core_http_response "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number,omitempty" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse UserResponseDto

func (h *UserHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponsehandler(log, w)

	var req CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}
	userDomain, err := h.userService.CreateUser(ctx, domainFromDTO(req))
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	responseHandler.JSONResponse(CreateUserResponse(userDtoFromDomain(userDomain)), http.StatusCreated)

}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUnitialized(dto.FullName, dto.PhoneNumber)
}
