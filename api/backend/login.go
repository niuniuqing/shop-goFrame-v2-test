package backend

import "github.com/gogf/gf/v2/frame/g"

type LoginDoReq struct {
	g.Meta   `path:"/backend/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}
type LoginDoRes struct {
	//todo
	//Referer string `json:"referer" dc:"引导客户端跳转地址"`
	//返回一个登录信息
	User interface{} `json:"user"`
}
