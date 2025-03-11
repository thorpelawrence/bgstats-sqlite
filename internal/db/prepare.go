package db

func (db *DB) prepareStmts() error {
	var err error

	db.insertPlayer, err = db.Prepare(`
		INSERT OR REPLACE INTO players (id, name, mtime)
		VALUES (?, ?, ?)
	`)
	if err != nil {
		return err
	}

	db.insertLocation, err = db.Prepare(`
		INSERT OR REPLACE INTO locations (id, name, mtime)
		VALUES (?, ?, ?)
	`)
	if err != nil {
		return err
	}

	return nil
}
