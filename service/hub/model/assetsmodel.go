package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AssetsModel = (*customAssetsModel)(nil)

type (
	// AssetsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAssetsModel.
	AssetsModel interface {
		assetsModel
	}

	customAssetsModel struct {
		*defaultAssetsModel
	}
)

// NewAssetsModel returns a model for the database table.
func NewAssetsModel(conn sqlx.SqlConn) AssetsModel {
	return &customAssetsModel{
		defaultAssetsModel: newAssetsModel(conn),
	}
}
