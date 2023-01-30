package position

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

type sPosition struct{}

func init() {
	service.RegisterPosition(New())
}

func New() *sPosition {
	return &sPosition{}
}

// 添加轮播图
func (s *sPosition) Create(ctx context.Context, in model.PositionCreateInput) (out model.PositionCreateOutput, err error) {

	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	lastInsertID, err := dao.PositionInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	//todo PositionId: int
	return model.PositionCreateOutput{PositionId: int(lastInsertID)}, err
}

// Delete 删除
func (s *sPosition) Delete(ctx context.Context, id uint) error {
	//原先是tx *gdb.TX，现在为tx gdb.TX
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除轮播图
		_, err := dao.PositionInfo.Ctx(ctx).Where(g.Map{
			dao.PositionInfo.Columns().Id: id,
			//}).Unscoped().Delete()   //Unscoped()物理删除一条数据
		}).Delete() //软删除数据，即只在删除时间列加上删除时间
		return err
	})
}

// Update 修改
func (s *sPosition) Update(ctx context.Context, in model.PositionUpdateInput) error {
	return dao.PositionInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}
		//链式操作
		_, err := dao.PositionInfo.
			Ctx(ctx).
			Data(in).
			//FieldsEx字段过滤
			FieldsEx(dao.PositionInfo.Columns().Id).
			Where(dao.PositionInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sPosition) GetList(ctx context.Context, in model.PositionGetListInput) (out *model.PositionGetListOutput, err error) {
	var (
		m = dao.PositionInfo.Ctx(ctx)
	)
	out = &model.PositionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.PositionInfo.Columns().Sort)

	// 执行查询
	var list []*entity.PositionInfo
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
	//if err := listModel.ScanList(&out.List, "Position"); err != nil {
	//查询单条模型数据比较简单，直接使用Scan方法即可
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
