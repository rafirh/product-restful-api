package web

type ProductRequestUpdate struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}
