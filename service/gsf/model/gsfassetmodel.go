package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GsfAssetModel = (*customGsfAssetModel)(nil)

type (
	// GsfAssetModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGsfAssetModel.
	GsfAssetModel interface {
		gsfAssetModel
	}

	customGsfAssetModel struct {
		*defaultGsfAssetModel
	}
)

// NewGsfAssetModel returns a model for the database table.
func NewGsfAssetModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) GsfAssetModel {
	return &customGsfAssetModel{
		defaultGsfAssetModel: newGsfAssetModel(conn, c, opts...),
	}
}
