package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/thorpelawrence/bgstats-sqlite/internal/export"
)

func (db *DB) ImportModel(model *export.Model) error {
	tx, err := db.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("beginning transaction: %w", err)
	}
	defer tx.Rollback()

	for _, player := range model.Players {
		if _, err := tx.Stmt(db.insertPlayer).Exec(player.ID, player.Name, player.Modified.Time); err != nil {
			return fmt.Errorf("creating player: %w", err)
		}
	}

	for _, location := range model.Locations {
		if _, err := tx.Stmt(db.insertLocation).Exec(location.ID, location.Name, location.Modified.Time); err != nil {
			return fmt.Errorf("creating location: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}
