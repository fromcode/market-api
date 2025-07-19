package sqlite

import (
	"database/sql"

	"github.com/fromcode/market-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		types TEXT,
		size INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateProduct(name string, types string, size int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO products (name, types, size) VALUES(?, ?, ?)")
	if err != nil {
		return 0, nil
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, types, size)
	if err != nil {
		return 0, nil
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return lastId, nil
}
