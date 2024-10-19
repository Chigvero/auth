package env

import (
	"errors"
	"github.com/Chigvero/auth/internal/config"
	"os"
)

var _ config.PGConfig = (*pgConfig)(nil)

const (
	dsnEnvName = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

func NewpgConfig() (*pgConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, errors.New("invalid dsn string")
	}
	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
