package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"tc-demo/database/fixtures/model"
	"time"
)

func WaitAndPingDBWithRetry(db *sql.DB, maxRetry int) error {
	var err error
	for tries := 0; tries <= maxRetry; tries++ {
		err = db.Ping()
		if err == nil {
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.Errorf("could not ping database: %v", err)
}

func MigrateUP(db *sql.DB, source string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(source, "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}
	return nil
}

func LoadProductFixtures(db *sql.DB, fixture string) error {
	fixtureData, err := ioutil.ReadFile(fixture)
	if err != nil {
		return errors.Errorf("could not read fixture data: %v", err)
	}

	// Unmarshal the YAML data into a slice of Product structs
	var products []model.Product
	err = yaml.Unmarshal(fixtureData, &products)
	if err != nil {
		return errors.Errorf("could not unmarshal fixture data: %v", err)
	}

	for _, product := range products {
		_, err = db.Exec(product.ToSQL())
		if err != nil {
			return errors.Errorf("could not execute SQL INSERT statement: %v", errors.Wrap(err, product.ToSQL()))
		}
	}
	return nil
}
