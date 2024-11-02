package env

import (
	"errors"
	"github.com/Chigvero/auth/internal/config"
	"net"
	"os"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

type grpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (*grpcConfig, error) {
	port := os.Getenv("GRPC_PORT")
	if len(port) == 0 {
		return nil, errors.New("Invalid port")
	}
	host := os.Getenv("GRPC_HOST")
	if len(host) == 0 {
		return nil, errors.New("Invalid host")
	}
	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (g *grpcConfig) Address() string {
	return net.JoinHostPort(g.host, g.port)
}
