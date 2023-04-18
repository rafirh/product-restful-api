package domain

import (
	"time"

)

type Product struct {
	Id int64
	Name string
	Price string
	CreatedAt time.Time
	UpdatedAt time.Time
}