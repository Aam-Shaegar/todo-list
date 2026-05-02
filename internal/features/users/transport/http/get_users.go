package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Aam-Shaegar/todo-list/internal/core/logger"
	core_http_request "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/request"
	core_http_response "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/response"
)

type GetUsersResponse []UserResponseDto

func (h *UserHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponsehandler(log, w)
	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get 'limit', 'offset' query parameters",
		)
		return
	}
	userDomains, err := h.userService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get users",
		)
		return
	}
	response := GetUsersResponse(usersDtoFromDomains(userDomains))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_request.GetINtQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}
	offset, err := core_http_request.GetINtQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}
	return limit, offset, nil

}
