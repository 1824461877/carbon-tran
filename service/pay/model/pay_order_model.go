package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PayOrderModel = (*customPayOrderModel)(nil)

type (
	// PayOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPayOrderModel.
	PayOrderModel interface {
		payOrderModel
	}

	customPayOrderModel struct {
		*defaultPayOrderModel
	}
)

// NewPayOrderModel returns a model for the database table.
func NewPayOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PayOrderModel {
	return &customPayOrderModel{
		defaultPayOrderModel: newPayOrderModel(conn, c, opts...),
	}
}
