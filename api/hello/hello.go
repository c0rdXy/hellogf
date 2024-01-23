package hello

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
