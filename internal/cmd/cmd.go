package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"hellogf/internal/controller/hello"
	"hellogf/internal/controller/user"
)

func handler(req *ghttp.Request) {
	req.Response.Writefln("Hello GF!")
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 函数注册
			//s.BindHandler("POST:/he", handler)

			// 对象注册
			//hello := hello.NewHello()
			//s.BindHandler("/say", hello.SayHello)
			//s.BindObject("/user", user.New(), "Add")
			//s.BindObjectRest("/user", user.New())

			// 分组注册
			//s.Group("/api", func(group *ghttp.RouterGroup) {
			//	group.Middleware(ghttp.MiddlewareHandlerResponse)
			//	group.Group("/hello", func(group *ghttp.RouterGroup) {
			//		group.Bind(
			//			hello.NewHello(),
			//		)
			//	})
			//	group.Group("/user", func(group *ghttp.RouterGroup) {
			//		group.Bind(
			//			user.New(),
			//		)
			//	})
			//})

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewHello(),
					user.New(),
				)
			})

			s.SetServerRoot("resource/public")
			s.Run()
			return nil
		},
	}
)
