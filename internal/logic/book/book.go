package book

import (
	"context"
	"hellogf/internal/dao"
	"hellogf/internal/model/do"
	"hellogf/internal/model/entity"
	"hellogf/internal/service"
)

func init() {
	service.RegisterBook(&sBook{})
}

type sBook struct{}

func (s sBook) GetList(ctx context.Context) (books []entity.Book, err error) {
	//TODO implement me
	err = dao.Book.Ctx(ctx).Scan(&books)
	return
}

func (s sBook) Add(ctx context.Context, book do.Book) (err error) {
	//TODO implement me
	_, err = dao.Book.Ctx(ctx).Insert(book)
	return
}

func (s sBook) Edit(ctx context.Context, book do.Book) (err error) {
	//TODO implement me
	_, err = dao.Book.Ctx(ctx).Update("id", book.Id)
	return
}

func (s sBook) Del(ctx context.Context, book do.Book) (err error) {
	//TODO implement me
	_, err = dao.Book.Ctx(ctx).Delete("id", book.Id)
	return
}
