package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TradeOrderModel = (*customTradeOrderModel)(nil)

type (
	// TradeOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTradeOrderModel.
	TradeOrderModel interface {
		tradeOrderModel
	}

	customTradeOrderModel struct {
		*defaultTradeOrderModel
	}
)

// NewTradeOrderModel returns a model for the database table.
func NewTradeOrderModel(conn sqlx.SqlConn) TradeOrderModel {
	return &customTradeOrderModel{
		defaultTradeOrderModel: newTradeOrderModel(conn),
	}
}
