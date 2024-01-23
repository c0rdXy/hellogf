// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package hello

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	hello2 "hellogf/api/hello"
)

type hello struct{}

func NewHello() *hello {
	return &hello{}
}

func (c *hello) SayHello(req *ghttp.Request) {
	req.Response.Writefln("你好GoFrame!")
}

func (c *hello) Params(ctx context.Context, req *hello2.ParamsReq) (res *hello2.ParamsRes, err error) {
	r := g.RequestFromCtx(ctx)
	/*	name := r.GetQuery("name", "李四")
		r.Response.Writeln(name.String() + "你好!")
		data := r.GetQueryMap()
		r.Response.Writeln(data)
		data := r.GetQueryMap(g.Map{"name": "李四", "age": 20})
		r.Response.Writeln(data)

		data := r.GetForm("username")

		data := r.GetFormMap()
		r.Response.Writeln(data)

		type user struct {
			UserName string
			Age      int
			Password string
		}

		var u user
		err = r.ParseForm(&u)

		name := r.GetRouter("name")

		name := r.GetRouterMap()

		data := r.Parse(&u)
		if err == nil {
			r.Response.Writeln(data)
		}*/

	r.Response.Writeln(req)

	return
}

func (c *hello) Response(ctx context.Context, req *hello2.ParamsReq) (res *hello2.ParamsRes, err error) {
	//r := g.RequestFromCtx(ctx)

	//r.Response.Writef("<h1>hello %s, age %d</h1>", req.UserName, req.Age)
	//r.Response.WriteExit("<h2 style='color: red'>hello</h2>")
	//r.Response.Writeln("<h2 style='color: red'>hello</h2>")
	//r.Response.Writeln("<h2 style='color: red'>hello</h2>")

	//r.Response.WriteJson(req)

	res = &hello2.ParamsRes{
		UserName: "www",
		Password: "123456",
		Age:      18,
	}

	err = gerror.New("服务器开小差了")

	return
}

func (c *hello) Db(req *ghttp.Request) {
	//md := g.Model("book")
	//book, err := md.One()
	//books, err := md.All()
	//books, err := md.Fields("id, name, price").All()
	//books, err := md.FieldsEx("id").All()
	//books, err := md.Array("name")
	min, err := g.Model("book").Min("price")
	max, err := g.Model("book").Max("price")
	avg, err := g.Model("book").Avg("price")
	count, err := g.Model("book").Count()
	if err == nil {
		req.Response.Writeln(g.Map{
			"min":   min,
			"max":   max,
			"avg":   avg,
			"count": count,
		})
	}
}

func (c *hello) Db1(req *ghttp.Request) {
	md := g.Model("book")
	//books, err := md.Where("id", 2).All()
	//books, err := md.WhereGTE("id", 2).WhereLTE("id", 4).All()
	//books, err := md.WhereIn("id", g.Array{3, 4, 5}).All()
	//books, err := md.WhereIn("id", g.Array{3, 4, 5}).WhereOrLike("name", "%数据%").All()
	//books, err := md.Order("price", "ASC").Order("id", "DESC").All()
	//books, err := md.Group("name").All()
	//books, err := md.Limit(3, 4).All()
	books, err := md.Page(2, 3).All()
	if err == nil {
		req.Response.Writeln(books)
	}
}

func (c *hello) Db2(req *ghttp.Request) {
	type Book struct {
		BookId  uint        `json:"id" orm:"id"`
		Name    string      `json:"name"`
		Author  string      `json:"author"`
		Price   float64     `json:"price"`
		PubTime *gtime.Time `orm:"publish_time" json:"pubTime"`
	}
	//var book *Book
	var book []Book

	md := g.Model("book")
	//err := md.Scan(&book)

	//err := md.Where("id", 10).Scan(&book)
	err := md.Scan(&book)

	if err == nil {
		if book != nil {
			req.Response.Writeln(book)
		}
	} else {
		req.Response.Writeln(err.Error())
	}
}
