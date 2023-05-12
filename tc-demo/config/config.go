package config

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

const DatabasePassword = "s3cr3t"
const DatabaseUser = "postgres"
const DatabaseName = "tc-demo"

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DBName:   DatabaseName,
		User:     DatabaseUser,
		Password: DatabasePassword,
	}
}

func (dc *DatabaseConfig) Update(ctx context.Context, ctr testcontainers.Container) error {
	host, err := ctr.Host(ctx)

	if err != nil {
		return err
	}

	port, err := ctr.MappedPort(ctx, "5432/tcp")

	if err != nil {
		return err
	}
	dc.Host = host
	dc.Port = port.Int()
	return nil
}

func (dc *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dc.Host,
		dc.Port,
		dc.User,
		dc.Password,
		dc.DBName,
	)
}
