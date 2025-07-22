package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/fromcode/market-api/internal/config"
	"github.com/fromcode/market-api/internal/types"
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

func (s *Sqlite) GetProductsById(id int64) (types.Markets, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, types, size FROM products WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Markets{}, err
	}

	defer stmt.Close()

	var market types.Markets

	err = stmt.QueryRow(id).Scan(&market.Id, &market.Name, &market.Type, &market.Size)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Markets{}, fmt.Errorf("product tidak ditemukan %s", fmt.Sprint(id))
		}
		return types.Markets{}, fmt.Errorf("query error: %w", err)
	}

	return market, nil
}
