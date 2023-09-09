package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserWalletModel = (*customUserWalletModel)(nil)

type (
	// UserWalletModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserWalletModel.
	UserWalletModel interface {
		userWalletModel
	}

	customUserWalletModel struct {
		*defaultUserWalletModel
	}
)

// NewUserWalletModel returns a model for the database table.
func NewUserWalletModel(conn sqlx.SqlConn, opts ...cache.Option) UserWalletModel {
	return &customUserWalletModel{
		defaultUserWalletModel: newUserWalletModel(conn),
	}
}
