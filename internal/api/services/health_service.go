package services

import (
	"context"
	"github.com/aguerram/gtcth/internal/api/dto/response"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

type HealthService interface {
	HealthCheck(ctx context.Context) (*response.HealthCheckResponse, error)
}

type healthService struct {
	db *pgx.Conn
}

func NewHealthService(conn *pgx.Conn) HealthService {
	return &healthService{
		db: conn,
	}
}

func (h *healthService) HealthCheck(ctx context.Context) (*response.HealthCheckResponse, error) {
	err := h.db.Ping(ctx)
	res := response.NewHealthCheckResponse()
	isDbUp := true
	if err != nil {
		log.Error("DB is down")
		isDbUp = false
	}
	res.AddComponentStatus("Database", isDbUp)

	return res, nil
}
