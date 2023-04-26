package domain

import (
	"time"

)

type Product struct {
	Id int
	Name string
	Price int64
	CreatedAt time.Time
	UpdatedAt time.Time
}