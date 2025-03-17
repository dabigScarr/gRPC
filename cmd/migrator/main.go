package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	emptyValue = ""
)

func main() {
	var storagePath, migrationPath, migrationTable string

	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationPath, "migration-path", "", "path to migrations")
	flag.StringVar(&migrationTable, "migrations-table", "migration", "name of migrations")
	flag.Parse()

	if storagePath == emptyValue {
		panic("storage-path is required")
	}
	if migrationPath == emptyValue {
		panic("migrations-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}
		panic(err)
	}

	fmt.Println("migrations applied successfully")

}
