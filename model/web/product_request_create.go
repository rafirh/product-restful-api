package web

type ProductRequestCreate struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
