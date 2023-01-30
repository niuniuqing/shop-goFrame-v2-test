// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsOptionsInfoDao is the data access object for table goods_options_info.
type GoodsOptionsInfoDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns GoodsOptionsInfoColumns // columns contains all the column names of Table for convenient usage.
}

// GoodsOptionsInfoColumns defines and stores column names for table goods_options_info.
type GoodsOptionsInfoColumns struct {
	Id        string //
	GoodsId   string // 商品id
	PicUrl    string // 图片
	Name      string // 商品名称
	Price     string // 价格 单位分
	Stock     string // 库存
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// goodsOptionsInfoColumns holds the columns for table goods_options_info.
var goodsOptionsInfoColumns = GoodsOptionsInfoColumns{
	Id:        "id",
	GoodsId:   "goods_id",
	PicUrl:    "pic_url",
	Name:      "name",
	Price:     "price",
	Stock:     "stock",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewGoodsOptionsInfoDao creates and returns a new DAO object for table data access.
func NewGoodsOptionsInfoDao() *GoodsOptionsInfoDao {
	return &GoodsOptionsInfoDao{
		group:   "default",
		table:   "goods_options_info",
		columns: goodsOptionsInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoodsOptionsInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoodsOptionsInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoodsOptionsInfoDao) Columns() GoodsOptionsInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoodsOptionsInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoodsOptionsInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoodsOptionsInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
