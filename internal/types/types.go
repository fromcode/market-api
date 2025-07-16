package types

type Markets struct {
	Id   int
	Name string `validate:"required"`
	Type string `validate:"required"`
	Size int    `validate:"required"`
}
