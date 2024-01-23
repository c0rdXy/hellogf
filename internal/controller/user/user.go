package user

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hellogf/api/user"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// func Handler(ctx context.Context, req *Request) (res *Response, err error)

func (c *Controller) Add(ctx context.Context, req *user.AddReq) (res *user.AddRes, err error) {
	// 返回Json数据
	res = &user.AddRes{
		Name: "张三",
		Age:  18,
	}
	// 返回自定义数据
	//g.RequestFromCtx(ctx).Response.Writefln("Hello")
	return
}

func (c *Controller) Update(r *ghttp.Request) {
	r.Response.Writefln("更新用户")
}

func (c *Controller) Delete(r *ghttp.Request) {
	r.Response.Writefln("删除用户")
}

func (c *Controller) Get(r *ghttp.Request) {
	r.Response.Writefln("获取用户")
}

func (c *Controller) Put(r *ghttp.Request) {
	r.Response.Writefln("修改用户")
}

func (c *Controller) Post(r *ghttp.Request) {
	r.Response.Writefln("新增用户")
}
