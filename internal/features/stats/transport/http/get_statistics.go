package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Aam-Shaegar/todo-list/internal/core/domain"
	core_logger "github.com/Aam-Shaegar/todo-list/internal/core/logger"
	core_http_request "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/request"
	core_http_response "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/response"
)

type GetStatisticsResponse struct {
	TasksCreated               int      `json:"tasks_created"`
	TasksCompleted             int      `json:"tasks_completed"`
	TasksCompletedRate         *float64 `json:"tasks_completed_rate"`
	TasksAverageCompletionTime *string  `json:"tasks_average_completion_time"`
}

func (h *StatisticsHTTPHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponsehandler(log, w)
	userID, from, to, err := getUserIdFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get query params",
		)
		return
	}
	stats, err := h.statisticsService.GetStats(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get statistics",
		)
		return
	}

	response := toDTOFromDomain(stats)
	responseHandler.JSONResponse(
		response,
		http.StatusOK,
	)
}

func getUserIdFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	userID, err := core_http_request.GetINtQueryParam(r, "user_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}
	from, err := core_http_request.GetDateQueryParam(r, "from")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' query param: %w", err)
	}
	to, err := core_http_request.GetDateQueryParam(r, "to")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' query param: %w", err)
	}

	return userID, from, to, nil
}

func toDTOFromDomain(stats domain.Statistics) GetStatisticsResponse {
	var avgTime *string
	if stats.TasksAverageCompletionTime != nil {
		duration := stats.TasksAverageCompletionTime.String()
		avgTime = &duration
	}
	return GetStatisticsResponse{
		TasksCreated:               stats.TasksCreated,
		TasksCompleted:             stats.TasksCompleted,
		TasksCompletedRate:         stats.TasksCompletedRate,
		TasksAverageCompletionTime: avgTime,
	}
}
