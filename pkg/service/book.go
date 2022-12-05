package service

import (
	"context"
	"ginapis/pkg/server/http/bean"
	"github.com/bbdshow/bkit/typ"
	"time"
)

func (svc *Service) ListBook(ctx context.Context, in *bean.ListBookReq, out *typ.ListResp) error {
	var books []bean.Book
	for i := 0; i < 3; i++ {
		books = append(books, bean.Book{
			Id:        0,
			Name:      "",
			Author:    "",
			Price:     0,
			CreatedAt: time.Time{},
		})
	}
	out.Count = 3
	out.List = books
	return nil
}
