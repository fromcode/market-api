package types

type Markets struct {
	Id   int64  `json:"id"`
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
	Size int    `json:"size" validate:"required"`
}
