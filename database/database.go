package database

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"path/filepath"
	"runtime"

	// imported to register the postgres driver
	_ "github.com/lib/pq"
	"strings"
	// imported to register the postgres migration driver
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// imported to register the "file" source migration driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Sets up the config connections using the provided configuration
func NewDB(config *Config) (*sqlx.DB, error) {
	connStr := dbConnectionString(config)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %v", err)
	}
	return db, nil
}

// dbConnectionString builds a connection string suitable for the pgx
// Postgres driver, using the values of vars
func dbConnectionString(config *Config) string {
	vals := dbValues(config)
	var p []string
	for k, v := range vals {
		p = append(p, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(p, " ")
}

// Runs the migrations. u is the connection URL (e.g. postgres://...).
func dbMigrate(u string) error {
	migrationsDir := fmt.Sprintf("file://%s", dbMigrationsDir())
	fmt.Errorf(migrationsDir)
	m, err := migrate.New(migrationsDir, u)
	if err != nil {
		return fmt.Errorf("failed create migrate: %w", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to run migrate: %w", err)
	}
	srcErr, dbErr := m.Close()
	if srcErr != nil {
		return fmt.Errorf("migrate source error: %w", srcErr)
	}
	if dbErr != nil {
		return fmt.Errorf("migrate database error: %w", dbErr)
	}
	return nil
}

// Return the path on disk to the migrations.
func dbMigrationsDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	return filepath.Join(filepath.Dir(filename), "../../migrations")
}
