package service

import (
	"context"
	"hellogf/internal/model/do"
	"hellogf/internal/model/entity"
)

// 1.定义接口
type IBook interface {
	GetList(ctx context.Context) (books []entity.Book, err error)
	Add(ctx context.Context, book do.Book) (err error)
	Edit(ctx context.Context, book do.Book) (err error)
	Del(ctx context.Context, book do.Book) (err error)
}

// 2.定义接口变量
var localBook IBook

// 3.定义一个获取接口实例的函数
func Book() IBook {
	if localBook == nil {
		panic("IBook接口未实现或未注册")
	}
	return localBook
}

// 4.定义一个接口实现的注册方法
func RegisterBook(I IBook) {
	localBook = I
}
