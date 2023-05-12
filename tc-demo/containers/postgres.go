package containers

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"tc-demo/config"
)

func NewPostgresContainer(ctx context.Context, cfg *config.DatabaseConfig) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image: "postgres:13.3-alpine",
		Env: map[string]string{
			"POSTGRES_PASSWORD": cfg.Password,
			"POSTGRES_USER":     cfg.User,
			"POSTGRES_DB":       cfg.DBName,
		},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections"),
	}
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
