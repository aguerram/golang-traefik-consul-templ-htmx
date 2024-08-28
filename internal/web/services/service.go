package services

import "github.com/jackc/pgx/v5"

type WebService struct {
	UserService UserService
}

func NewService(db *pgx.Conn) *WebService {
	return &WebService{
		UserService: NewUserService(db),
	}
}
