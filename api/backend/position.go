package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// api层（业务侧）和model层（数据层）基本是一样的
// 手工位添加
type PositionReq struct {
	//访问路径
	g.Meta `path:"/backend/position/add" tags:"Position" method:"post" summary:"You first position api"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"       v:"required#跳转链接不能为空"      dc:"跳转链接"`
	//商品的名称是可以通过id取出来的，但是由于项目之后的一些问题，做一些冗余处理是有必要的
	GoodsName string `json:"goods_name"   v:"required#商品名称不能为空"      dc:"商品名称"` //冗余设计
	GoodsId   uint   `json:"goods_id"   v:"required#商品Id不能为空"      dc:"商品Id"`
	Sort      int    `json:"sort"       dc:"排序"`
}
type PositionRes struct {
	//g.Meta  `mime:"text/html" example:"string"`
	PositionId int `json:"positionId"`
}

// 手工位删除
type PositionDeleteReq struct {
	g.Meta `path:"/backend/position/delete" method:"delete" tags:"手工位" summary:"删除手工位接口"`
	Id     uint `v:"min:1#请选择需要删除的手工位" dc:"手工位id"`
}
type PositionDeleteRes struct{}

// 修改手工位
type PositionUpdateReq struct {
	g.Meta    `path:"/backend/position/update/{Id}" method:"post" tags:"手工位" summary:"修改手工位接口"`
	Id        uint   `json:"id"          v:"min:1#请选择需要修改的手工位" dc:"手工位Id"`
	PicUrl    string `json:"pic_url"     dc:"图片链接"`
	Link      string `json:"link"        dc:"跳转链接"`
	GoodsName string `json:"goods_name"  dc:"商品名称"`
	GoodsId   uint   `json:"goods_id"    dc:"商品Id"`
	Sort      int    `json:"sort"        dc:"图片排序"`
}
type PositionUpdateRes struct {
	Id uint `json:"id"`
}

// 取列表数据
type PositionGetListCommonReq struct {
	g.Meta `path:"/backend/position/list" method:"get" tags:"手工位" summary:"手工位列表接口"`
	Sort   int `json:"sort"   in:"query" dc:"排序类型"`
	CommonPaginationReq
}
type PositionGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
