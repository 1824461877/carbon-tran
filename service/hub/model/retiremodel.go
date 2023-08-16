package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RetireModel = (*customRetireModel)(nil)

type (
	// RetireModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRetireModel.
	RetireModel interface {
		retireModel
	}

	customRetireModel struct {
		*defaultRetireModel
	}
)

// NewRetireModel returns a model for the database table.
func NewRetireModel(conn sqlx.SqlConn) RetireModel {
	return &customRetireModel{
		defaultRetireModel: newRetireModel(conn),
	}
}
