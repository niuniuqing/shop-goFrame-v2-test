package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// api层（业务侧）和model层（数据层）基本是一样的
// 管理员添加
type AdminReq struct {
	//访问路径
	g.Meta   `path:"/backend/admin/add" tags:"Admin" method:"post" summary:"You first admin api"`
	Name     string `json:"name''"     v:"required#用户名接不能为空"   dc:"用户名"`
	Password string `json:"password"   v:"required#密码不能为空"      dc:"密码"`
	RoleIds  string `json:"role_ids"   dc:"角色ids，超管不需要"`
	IsAdmin  int    `json:"is_admin"   dc:"是否是超级管理员"`
}
type AdminRes struct {
	//g.Meta  `mime:"text/html" example:"string"`
	AdminId int `json:"admin_id"`
}

// 管理员删除
type AdminDeleteReq struct {
	g.Meta `path:"/backend/admin/delete" method:"delete" tags:"管理员" summary:"删除管理员接口"`
	Id     uint `v:"min:1#请选择需要删除的管理员" dc:"管理员id"`
}
type AdminDeleteRes struct{}

// 修改管理员
type AdminUpdateReq struct {
	g.Meta   `path:"/backend/admin/update/{Id}" method:"post" tags:"管理员" summary:"修改管理员接口"`
	Id       uint   `json:"id"         v:"min:1#请选择需要修改的管理员" dc:"管理员Id"`
	Name     string `json:"name''"     v:"required#用户名接不能为空"   dc:"用户名"`
	Password string `json:"password"   dc:"密码"`
	RoleIds  string `json:"role_ids"   dc:"角色ids，超管不需要"`
	IsAdmin  int    `json:"is_admin"   dc:"是否是超级管理员"`
}
type AdminUpdateRes struct {
	Id uint `json:"id"`
}

// 取列表数据
type AdminGetListCommonReq struct {
	g.Meta `path:"/backend/admin/list" method:"get" tags:"管理员" summary:"管理员列表接口"`
	CommonPaginationReq
}
type AdminGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
