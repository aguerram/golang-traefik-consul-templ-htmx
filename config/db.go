package config

import (
	"context"
	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"
)

func NewDatabaseConnection(env *AppEnv) (*pgx.Conn, func(ctx context.Context), error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, env.DSN)
	if err != nil {
		return nil, nil, err
	}
	close := func(ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			return
		}
	}
	log.Info("Connected to database")
	return conn, close, nil
}
