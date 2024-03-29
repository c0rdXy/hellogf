// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package hello

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	hello2 "hellogf/api/hello"
	"hellogf/internal/dao"
	"hellogf/internal/model/entity"
	"hellogf/internal/service"
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

	books, err := service.Book().GetList(req.Context())

	if err == nil {
		req.Response.WriteJson(books)
	} else {
		req.Response.WriteJson(err.Error())
	}
}

func (c *hello) Tpl(req *ghttp.Request) {
	html := `<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>

    <ul>
        <li>姓名：{{.name}}</li>
        <li>性别：{{.gender}}</li>
        <li>年龄：{{.age}}</li>
    </ul>

</body>
</html>`

	//req.Response.Writef(html, "张三", "男", 18)
	req.Response.WriteTplContent(html, g.Map{
		"name":   "张三",
		"gender": "男",
		"age":    18,
	})

}

func (c *hello) Tpl1(req *ghttp.Request) {
	var books []entity.Book
	dao.Book.Ctx(req.Context()).Scan(&books)

	req.Response.WriteTpl("/hello/index.html", g.Map{
		"slice": g.Slice{"张三", "李四", "王五"},
		"map": g.Map{
			"name":   "张三",
			"gender": "男",
			"age":    18,
		},
		"books": books,
		"num":   150,
	})

}

func (c *hello) Tpl2(req *ghttp.Request) {
	var books []entity.Book
	dao.Book.Ctx(req.Context()).Scan(&books)

	req.Response.WriteTpl("/hello/index1.html", g.Map{
		"books": books,
	})

}

func (c *hello) Upload(req *ghttp.Request) {
	file := req.GetUploadFile("ufile")

	if file != nil {
		//req.Response.Writeln(file)
		file.Filename = "test.png"
		filename, err := file.Save("resource/public/upload/")
		if err == nil {
			req.Response.Writeln("/upload/" + filename)
		} else {
			req.Response.Writeln(err)
		}
	}
}

func (c *hello) Download(req *ghttp.Request) {
	//req.Response.ServeFile("resource/public/upload/test.png")
	//req.Response.ServeFileDownload("resource/public/upload/test.png")
	req.Response.ServeFileDownload("resource/public/upload/test.png", "123.png")
}

func (c *hello) Valid(ctx context.Context, req *hello2.ValidReq) (res *hello2.ValidRes, err error) {

	return
}
