package storage

type Storage interface {
	CreateProduct(name string, types string, size int) (int64, error)
}
