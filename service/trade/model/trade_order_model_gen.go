// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tradeOrderFieldNames          = builder.RawFieldNames(&TradeOrder{})
	tradeOrderRows                = strings.Join(tradeOrderFieldNames, ",")
	tradeOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(tradeOrderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tradeOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(tradeOrderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tradeOrderModel interface {
		Insert(ctx context.Context, data *TradeOrder) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TradeOrder, error)
		FindOneByInitiatorTradeOrderIdList(ctx context.Context, initiator string) (*[]TradeOrder, error)
		FindOneByRecipientTradeOrderIdList(ctx context.Context, recipient string) (*[]TradeOrder, error)
		FindOneByTradeOrderId(ctx context.Context, tradeOrderId string) (*TradeOrder, error)
		FindOneByPayOrderId(ctx context.Context, payOrderId string) (*TradeOrder, error)
		Update(ctx context.Context, data *TradeOrder) error
		UpdateTradeStatus(ctx context.Context,PayOrderId string,TradeStatus int64) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTradeOrderModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TradeOrder struct {
		Id            int64     `db:"id"`
		TradeOrderId  string    `db:"trade_order_id"`  // 交易订单id
		PayOrderId    string    `db:"pay_order_id"`    // 支付订单id
		ExchangeAssetID string  `db:"exchange_asset_id"`  // 交易所订单 id
		CarbonAssetId string    `db:"carbon_asset_id"` // 用户密码
		CollectionID  string    `db:"collection_id"`   // 收款账号
		Initiator     string    `db:"initiator"`       // 交易的发起者
		Recipient     string    `db:"recipient"`       // 交易接受者
		TradeStatus   int64     `db:"trade_status"`    // 交易状态
		Number        int64     `db:"number"`          // 数量
		InitiatorTime time.Time `db:"initiator_time"`
		FinishTime    time.Time `db:"finish_time"`
	}
)

func newTradeOrderModel(conn sqlx.SqlConn) *defaultTradeOrderModel {
	return &defaultTradeOrderModel{
		conn:  conn,
		table: "`trade_order`",
	}
}

func (m *defaultTradeOrderModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTradeOrderModel) FindOne(ctx context.Context, id int64) (*TradeOrder, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tradeOrderRows, m.table)
	var resp TradeOrder
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTradeOrderModel) FindOneByTradeOrderId(ctx context.Context, tradeOrderId string) (*TradeOrder, error) {
	query := fmt.Sprintf("select %s from %s where `trade_order_id` = ? limit 1", tradeOrderRows, m.table)
	var resp TradeOrder
	err := m.conn.QueryRowCtx(ctx, &resp, query, tradeOrderId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTradeOrderModel) FindOneByPayOrderId(ctx context.Context, payOrderId string) (*TradeOrder, error) {
	query := fmt.Sprintf("select %s from %s where `pay_order_id` = ? limit 1", tradeOrderRows, m.table)
	var resp TradeOrder
	err := m.conn.QueryRowCtx(ctx, &resp, query, payOrderId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}


func (m *defaultTradeOrderModel) FindOneByInitiatorTradeOrderIdList(ctx context.Context, initiator string) (*[]TradeOrder, error) {
	var resp []TradeOrder
	query := fmt.Sprintf("select %s from %s where `initiator` = ?", tradeOrderRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, initiator)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTradeOrderModel) FindOneByRecipientTradeOrderIdList(ctx context.Context, recipient string) (*[]TradeOrder, error) {
	var resp []TradeOrder
	query := fmt.Sprintf("select %s from %s where `recipient` = ?", tradeOrderRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, recipient)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}


func (m *defaultTradeOrderModel) Insert(ctx context.Context, data *TradeOrder) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tradeOrderRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.TradeOrderId, data.PayOrderId, data.ExchangeAssetID, data.CarbonAssetId, data.CollectionID, data.Initiator, data.Recipient, data.TradeStatus, data.Number, data.InitiatorTime, data.FinishTime)
	return ret, err
}

func (m *defaultTradeOrderModel) Update(ctx context.Context, newData *TradeOrder) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tradeOrderRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.TradeOrderId, newData.ExchangeAssetID,  newData.PayOrderId, newData.CarbonAssetId, newData.CollectionID, newData.Initiator, newData.Recipient, newData.TradeStatus, newData.Number, newData.InitiatorTime, newData.FinishTime, newData.Id)
	return err
}

func (m *defaultTradeOrderModel) UpdateTradeStatus(ctx context.Context,PayOrderId string,TradeStatus int64) error {
	query := fmt.Sprintf("update %s set `trade_status`= ? where `pay_order_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, TradeStatus, PayOrderId)
	return err
}

func (m *defaultTradeOrderModel) tableName() string {
	return m.table
}
