package users_transport_http

import (
	"net/http"

	core_logger "github.com/Aam-Shaegar/todo-list/internal/core/logger"
	core_http_request "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/request"
	core_http_response "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/response"
)

func (h *UserHTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponsehandler(logger, w)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get ID from path value")
		return
	}
	if err := h.userService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}
	responseHandler.NoContentResponse()
}
