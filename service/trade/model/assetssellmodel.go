package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AssetsSellModel = (*customAssetsSellModel)(nil)

type (
	// AssetsSellModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAssetsSellModel.
	AssetsSellModel interface {
		assetsSellModel
	}

	customAssetsSellModel struct {
		*defaultAssetsSellModel
	}
)

// NewAssetsSellModel returns a model for the database table.
func NewAssetsSellModel(conn sqlx.SqlConn, c cache.CacheConf) AssetsSellModel {
	return &customAssetsSellModel{
		defaultAssetsSellModel: newAssetsSellModel(conn),
	}
}
