package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

const DB_NAME = "social.db"

func CreateDB() (*SQLiteRepository, error) {
	// create social.db if not exists
	if _, err := os.Stat(DB_NAME); os.IsNotExist(err) {
		_, err = os.Create(DB_NAME)
		if err != nil {
			fmt.Println("Cannot create db file")
			return nil, err
		}
		fmt.Println("Created db file")
	}

	// open database with foreign keys on
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=on", DB_NAME))
	if err != nil {
		fmt.Println("Cannot open db")
		return nil, err
	}

	// create sqilte driver
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		fmt.Println("Cannot create driver")
		return nil, err
	}

	// migration
	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migrations/",
		"sqlite3", driver)
	if err != nil {
		fmt.Println("Cannot create instance")
		return nil, err
	}

	// migration down
	err = m.Down()
	if err != nil {
		fmt.Println("Cannot migrate down")
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		}
	}

	// migration up
	err = m.Up()
	if err != nil {
		fmt.Println("Cannot migrate up")
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		}
	}

	return NewSQLiteRepository(db), nil
}
