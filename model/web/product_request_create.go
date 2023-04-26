package web

type ProductRequestCreate struct {
	Name  string `validate:"required,min=1,max=255"`
	Price int64    `validate:"required,min=0"`
}
