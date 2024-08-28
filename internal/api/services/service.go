package services

import "github.com/jackc/pgx/v5"

type ApiService struct {
	HealthService HealthService
}

func NewService(conn *pgx.Conn) *ApiService {
	return &ApiService{
		HealthService: NewHealthService(conn),
	}
}
