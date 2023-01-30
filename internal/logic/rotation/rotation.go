package rotation

import (
	"context"

	"github.com/gogf/gf/v2/encoding/ghtml"
	"shop-goFrame-v2-test/internal/service"

	"shop-goFrame-v2-test/internal/dao"
	"shop-goFrame-v2-test/internal/model"
)

type sRotation struct{}

func init() {
	service.RegisterRotation(New())
}

func New() *sRotation {
	return &sRotation{}
}

// 添加轮播图
func (s *sRotation) Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.RotationInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	//todo RotationId: int
	return model.RotationCreateOutput{RotationId: int(lastInsertID)}, err
}
