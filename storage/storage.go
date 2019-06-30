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

// Settings holds all storage configuration settings
type Settings struct {
	Type              string `mapstructure:"type" description:"Type of the storage"`
	URI               string `mapstructure:"uri" description:"URI of the storage"`
	SkipSSLValidation bool   `mapstructure:"skip_ssl_validation" description:"whether to skip ssl verification when connecting to the storage"`
}

// DefaultSettings returns default settings for the storage
func DefaultSettings() *Settings {
	return &Settings{
		Type:              "",
		URI:               "",
		SkipSSLValidation: false,
	}
}

// Validate validates the storage settings
func (s *Settings) Validate() error {
	if len(s.Type) == 0 {
		return fmt.Errorf("validate Settings: StorageType missing")
	}
	if len(s.URI) == 0 {
		return fmt.Errorf("validate Settings: StorageURI missing")
	}
	return nil
}

// Storage interface represents a Transactional Database
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
func New(s *Settings) (Storage, error) {
	if s.SkipSSLValidation {
		s.URI = s.URI + "?sslmode=disable"
	}

	db, err := sql.Open(s.Type, s.URI)
	if err != nil {
		return nil, fmt.Errorf("unable to open db connection: %s", err)
	}

	err = updateSchema(db, s.Type)
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
