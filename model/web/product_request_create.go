package web

type ProductRequestCreate struct {
	Name  string `validate:"required,min=1,max=255" json:"name"`
	Price int64 `validate:"required" json:"price"`
}
