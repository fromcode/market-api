package types

type Markets struct {
	Id   int64
	Name string `validate:"required"`
	Type string `validate:"required"`
	Size int    `validate:"required"`
}
