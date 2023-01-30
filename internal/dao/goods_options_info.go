// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"shop-goFrame-v2-test/internal/dao/internal"
)

// internalGoodsOptionsInfoDao is internal type for wrapping internal DAO implements.
type internalGoodsOptionsInfoDao = *internal.GoodsOptionsInfoDao

// goodsOptionsInfoDao is the data access object for table goods_options_info.
// You can define custom methods on it to extend its functionality as you wish.
type goodsOptionsInfoDao struct {
	internalGoodsOptionsInfoDao
}

var (
	// GoodsOptionsInfo is globally public accessible object for table goods_options_info operations.
	GoodsOptionsInfo = goodsOptionsInfoDao{
		internal.NewGoodsOptionsInfoDao(),
	}
)

// Fill with you ideas below.
