package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thorpelawrence/bgstats-sqlite/internal/db"
	"github.com/thorpelawrence/bgstats-sqlite/internal/export"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "modernc.org/sqlite"
)

var (
	out = flag.String("out", "bgstats.sqlite", "output sqlite db file")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error encountered:", err)
		os.Exit(1)
	}
}

func run() error {
	db, err := db.New(filepath.Clean(*out))
	if err != nil {
		return fmt.Errorf("running migrations: %w", err)
	}

	in := flag.Arg(0)
	if in == "" {
		// TODO check file presence
		return fmt.Errorf("no input file provided")
	}

	model, err := export.LoadFile(in)
	if err != nil {
		return fmt.Errorf("loading data model: %w", err)
	}

	if err := db.ImportModel(model); err != nil {
		return fmt.Errorf("importing model: %w", err)
	}

	return nil
}
