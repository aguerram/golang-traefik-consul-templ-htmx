package services

import "github.com/jackc/pgx/v5"

type Service struct {
	HealthService HealthService
}

func NewService(conn *pgx.Conn) *Service {
	return &Service{
		HealthService: NewHealthService(conn),
	}
}
