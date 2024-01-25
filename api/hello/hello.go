package hello

import "github.com/gogf/gf/v2/frame/g"

type ParamsReq struct {
	//g.Meta `path:"/params" method:"All"`

	UserName string `p:"name" d:"张三"`
	Password string
	Age      int
}

type ParamsRes struct {
	UserName string `p:"name" d:"张三" json:"user_name"`
	Password string `json:"pwd"`
	Age      int    `json:"age"`
}

type ValidReq struct {
	g.Meta `method:"All"`

	UserName string `p:"user_name" v:"required|length:2,16#请输入用户名|用户名长度应当在:min到:max之间"`
	Password string `p:"password" v:"required|password"`
	Age      int    `p:"age" v:"required|between:1,100#请输入年龄|年龄应当在:1到:100之间"`
}

type ValidRes struct{}
