package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/golang-migrate/migrate"
	migratepg "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

var (
	_, file, _, _ = runtime.Caller(0)
	basepath      = path.Dir(file)
)

// Config holds all storage configuration settings
type Config struct {
	Type string
	URI  string
}

type Storage interface {
	Get() *sql.DB
	Transaction(context context.Context, operation func(context context.Context, tx *sql.Tx) error) error
	Close() error
}

type postgresStorage struct {
	db *sql.DB
}

func (ps *postgresStorage) Get() *sql.DB {
	return ps.db
}

func (ps *postgresStorage) Transaction(context context.Context, operation func(context context.Context, tx *sql.Tx) error) error {
	tx, err := ps.db.BeginTx(context, nil)
	if err != nil {
		return err
	}

	if opErr := operation(context, tx); opErr != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return opErr
	}
	return tx.Commit()
}

func (ps *postgresStorage) Close() error {
	return ps.db.Close()
}

// New creates a Storage object and updates the database with the latest migrations
func New() (Storage, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:1234/postgres?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("unable to open db connection: %s", err)
	}

	err = updateSchema(db, "postgres")
	if err != nil {
		return nil, fmt.Errorf("unable to update db schema: %s", err)
	}

	log.Println("Database is up-to-date")

	return &postgresStorage{
		db: db,
	}, nil
}

func updateSchema(db *sql.DB, dbType string) error {
	driver, err := migratepg.WithInstance(db, &migratepg.Config{})
	if err != nil {
		return err
	}

	migrationsURL := fmt.Sprintf("file://%s/migrations", basepath)
	m, err := migrate.NewWithDatabaseInstance(migrationsURL, dbType, driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err == migrate.ErrNoChange {
		log.Println("No changes to the database schema have been made")
		return nil
	}
	return err
}
