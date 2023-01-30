package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

// api层（业务侧）和model层（数据层）基本是一样的
type RotationReq struct {
	//访问路径
	g.Meta `path:"/backend/rotation/add" tags:"Rotation" method:"post" summary:"You first rotation api"`
	PicUrl string `json:"pic_url"    v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"       v:"required#跳转链接不能为空"      dc:"跳转链接"`
	Sort   int    `json:"sort"       dc:"排序"`
}
type RotationRes struct {
	//todo
	//g.Meta  `mime:"text/html" example:"string"`
	RotationId int `json:"rotationId"`
}
