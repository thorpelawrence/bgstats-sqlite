package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/thorpelawrence/bgstats-sqlite/migrations"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB

	insertPlayer   *sql.Stmt
	insertLocation *sql.Stmt
}

func New(path string) (*DB, error) {
	d, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("opening sqlite db: %w", err)
	}

	db := &DB{DB: d}

	if err := db.migrateUp(); err != nil {
		return nil, fmt.Errorf("running migrations: %w", err)
	}

	if err := db.prepareStmts(); err != nil {
		return nil, fmt.Errorf("preparing statements: %w", err)
	}

	return db, nil
}

func (db *DB) migrateUp() error {
	migrationsIOFS, err := iofs.New(migrations.FS, ".")
	if err != nil {
		return fmt.Errorf("opening fs: %w", err)
	}

	instance, err := sqlite.WithInstance(db.DB, &sqlite.Config{})
	if err != nil {
		return fmt.Errorf("getting sqlite instance: %w", err)
	}
	m, err := migrate.NewWithInstance("iofs", migrationsIOFS, "sqlite", instance)
	if err != nil {
		return fmt.Errorf("loading migrations: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("running migrate up: %w", err)
	}

	return nil
}
