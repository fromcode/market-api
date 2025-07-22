package storage

import "github.com/fromcode/market-api/internal/types"

type Storage interface {
	CreateProduct(name string, types string, size int) (int64, error)
	GetProductsById(id int64) (types.Markets, error)
}
