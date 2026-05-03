package statistics_transport_http

import (
	"context"
	"net/http"
	"time"

	"github.com/Aam-Shaegar/todo-list/internal/core/domain"
	core_http_server "github.com/Aam-Shaegar/todo-list/internal/core/transport/http/server"
)

type StatisticsHTTPHandler struct {
	statisticsService StatisticsService
}

type StatisticsService interface {
	GetStats(ctx context.Context, userId *int, from, to *time.Time) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(statisticsService StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statisticsService: statisticsService,
	}
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStats,
		},
	}
}
