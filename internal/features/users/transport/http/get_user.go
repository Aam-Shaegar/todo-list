package users_transport_http

import (
	"net/http"

	core_logger "github.com/Aam-Shaegar/todo-list/internal/core/logger"
	core_http_response "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/response"
	core_http_utils "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/utils"
)

type GetUserResponse UserResponseDto

func (h *UserHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponsehandler(log, w)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get ID from path value")
	}
	user, err := h.userService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}
	response := GetUserResponse(userDtoFromDomain(user))
	responseHandler.JSONResponse(response, http.StatusOK)
}
