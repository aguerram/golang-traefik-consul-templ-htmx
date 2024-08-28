package services

import (
	"context"
	"github.com/aguerram/gtcth/internal/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	log "github.com/sirupsen/logrus"
	"time"
)

type UserService interface {
	ListUsers(ctx context.Context) ([]db.User, error)
}

type userServiceImpl struct {
	db *pgx.Conn
}

func NewUserService(db *pgx.Conn) UserService {
	return &userServiceImpl{
		db: db,
	}
}

func (u *userServiceImpl) ListUsers(ctx context.Context) ([]db.User, error) {
	log.Info("Attempt to list users")
	queries := db.New(u.db)
	var users, err = queries.GetAllUsers(ctx, db.GetAllUsersParams{
		StartDate: pgtype.Timestamp{
			Time:  time.Now().Add(-time.Hour * 24),
			Valid: true,
		},
		EndDate: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		StartOffset: 0,
		MaxElements: 10,
	})
	if err != nil {
		log.Errorf("Failed to list users: %v", err)
		return nil, err
	}
	log.Infof("Successfully listed %d users", len(users))
	return users, nil
}
