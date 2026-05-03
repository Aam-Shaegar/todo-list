package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Aam-Shaegar/todo-list/internal/core/logger"
	core_http_request "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/request"
	core_http_response "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResponse

func (h *TasksHTTPHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponsehandler(log, w)

	userID, limit, offset, err := getUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get query params",
		)
		return
	}

	tasksDomains, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get tasks",
		)
		return
	}
	response := GetTasksResponse(taskDTOsFromDomains(tasksDomains))
	responseHandler.JSONResponse(
		response,
		http.StatusOK,
	)
}

func getUserIDLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	limit, err := core_http_request.GetINtQueryParam(r, "limit")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}
	offset, err := core_http_request.GetINtQueryParam(r, "offset")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}
	userID, err := core_http_request.GetINtQueryParam(r, "user_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}
	return userID, limit, offset, nil

}
