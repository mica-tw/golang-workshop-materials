package cmd

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"gorm-demo/internal/config"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run the database migrations",
}

var upMigrateCmd = &cobra.Command{
	Use:   "up",
	Short: "Run up migration scripts",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()

		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		db, err := sql.Open("postgres", cfg.Database.DSN())
		if err != nil {
			return fmt.Errorf("failed to open postgres connection: %v", err)
		}
		defer db.Close()

		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return fmt.Errorf("failed to create postgres driver: %v", err)
		}

		m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
		if err != nil {
			return fmt.Errorf("failed to create migration instance: %v", err)
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("failed to apply migrations: %v", err)
		}

		fmt.Println("Database up migration successful")

		return nil
	},
}

var downMigrateCmd = &cobra.Command{
	Use:   "down",
	Short: "Run down migration scripts",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()

		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		db, err := sql.Open("postgres", cfg.Database.DSN())
		if err != nil {
			return fmt.Errorf("failed to open postgres connection: %v", err)
		}
		defer db.Close()

		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return fmt.Errorf("failed to create postgres driver: %v", err)
		}

		m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
		if err != nil {
			return fmt.Errorf("failed to create migration instance: %v", err)
		}

		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("failed to apply migrations: %v", err)
		}

		fmt.Println("Database down migration successful")

		return nil
	},
}

func init() {
	migrateCmd.AddCommand(upMigrateCmd)
	migrateCmd.AddCommand(downMigrateCmd)
	rootCmd.AddCommand(migrateCmd)
}
