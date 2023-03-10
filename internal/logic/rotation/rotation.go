package rotation

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"shop-goFrame-v2-test/internal/model/entity"

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

// Delete 删除
func (s *sRotation) Delete(ctx context.Context, id uint) error {
	//原先是tx *gdb.TX，现在为tx gdb.TX
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除轮播图
		_, err := dao.RotationInfo.Ctx(ctx).Where(g.Map{
			dao.RotationInfo.Columns().Id: id,
			//}).Unscoped().Delete()   //Unscoped()物理删除一条数据
		}).Delete() //软删除数据，即只在删除时间列加上删除时间
		return err
	})
}

// Update 修改
func (s *sRotation) Update(ctx context.Context, in model.RotationUpdateInput) error {
	return dao.RotationInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		//链式操作
		_, err := dao.RotationInfo.
			Ctx(ctx).
			Data(in).
			//FieldsEx字段过滤
			FieldsEx(dao.RotationInfo.Columns().Id).
			Where(dao.RotationInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询轮播图列表
func (s *sRotation) GetList(ctx context.Context, in model.RotationGetListInput) (out *model.RotationGetListOutput, err error) {
	var (
		m = dao.RotationInfo.Ctx(ctx)
	)
	out = &model.RotationGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.RotationInfo.Columns().Sort)

	// 执行查询
	var list []*entity.RotationInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	//查询多条数据记录并绑定数据到数据模型数组中，需要使用到ScanList方法
	//if err := listModel.ScanList(&out.List, "Rotation"); err != nil {
	//查询单条模型数据比较简单，直接使用Scan方法即可
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
