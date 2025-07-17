package storage

type Storage interface {
	CreateProduct(Name string, Type string, size int) (int64, error)
}
