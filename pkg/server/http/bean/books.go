package bean

import (
	"github.com/bbdshow/bkit/typ"
	"time"
)

type ListBookReq struct {
	typ.PageReq
	Name   string `form:"name" binding:"omitempty"`
	Author string `form:"name"`
}

type Book struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Author    string    `json:"author"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}
