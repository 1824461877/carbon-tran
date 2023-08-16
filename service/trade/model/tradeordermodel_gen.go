// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tradeOrderFieldNames          = builder.RawFieldNames(&TradeOrder{})
	tradeOrderRows                = strings.Join(tradeOrderFieldNames, ",")
	tradeOrderRowsExpectAutoSet   = strings.Join(stringx.Remove(tradeOrderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tradeOrderRowsWithPlaceHolder = strings.Join(stringx.Remove(tradeOrderFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTradeOrderIdPrefix           = "cache:tradeOrder:id:"
	cacheTradeOrderTradeOrderIdPrefix = "cache:tradeOrder:tradeOrderId:"
)

type (
	tradeOrderModel interface {
		Insert(ctx context.Context, data *TradeOrder) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TradeOrder, error)
		FindOneByTradeOrderId(ctx context.Context, tradeOrderId string) (*TradeOrder, error)
		Update(ctx context.Context, data *TradeOrder) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTradeOrderModel struct {
		sqlc.CachedConn
		table string
	}

	TradeOrder struct {
		Id            int64     `db:"id"`
		TradeOrderId  string    `db:"trade_order_id"`  // 交易订单id
		PayOrderId    string    `db:"pay_order_id"`    // 支付订单id
		CarbonAssetId string    `db:"carbon_asset_id"` // 用户密码
		Initiator     string    `db:"initiator"`       // 交易的发起者
		Recipient     string    `db:"recipient"`       // 交易接受者
		TadeStatus    int64     `db:"tade_status"`     // 交易状态
		InitiatorTime time.Time `db:"initiator_time"`
		FinishTime    time.Time `db:"finish_time"`
	}
)

func newTradeOrderModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTradeOrderModel {
	return &defaultTradeOrderModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`trade_order`",
	}
}

func (m *defaultTradeOrderModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	tradeOrderIdKey := fmt.Sprintf("%s%v", cacheTradeOrderIdPrefix, id)
	tradeOrderTradeOrderIdKey := fmt.Sprintf("%s%v", cacheTradeOrderTradeOrderIdPrefix, data.TradeOrderId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tradeOrderIdKey, tradeOrderTradeOrderIdKey)
	return err
}

func (m *defaultTradeOrderModel) FindOne(ctx context.Context, id int64) (*TradeOrder, error) {
	tradeOrderIdKey := fmt.Sprintf("%s%v", cacheTradeOrderIdPrefix, id)
	var resp TradeOrder
	err := m.QueryRowCtx(ctx, &resp, tradeOrderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tradeOrderRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
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
	tradeOrderTradeOrderIdKey := fmt.Sprintf("%s%v", cacheTradeOrderTradeOrderIdPrefix, tradeOrderId)
	var resp TradeOrder
	err := m.QueryRowIndexCtx(ctx, &resp, tradeOrderTradeOrderIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `trade_order_id` = ? limit 1", tradeOrderRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, tradeOrderId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
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
	tradeOrderIdKey := fmt.Sprintf("%s%v", cacheTradeOrderIdPrefix, data.Id)
	tradeOrderTradeOrderIdKey := fmt.Sprintf("%s%v", cacheTradeOrderTradeOrderIdPrefix, data.TradeOrderId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, tradeOrderRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.TradeOrderId, data.PayOrderId, data.CarbonAssetId, data.Initiator, data.Recipient, data.TadeStatus, data.InitiatorTime, data.FinishTime)
	}, tradeOrderIdKey, tradeOrderTradeOrderIdKey)
	return ret, err
}

func (m *defaultTradeOrderModel) Update(ctx context.Context, newData *TradeOrder) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	tradeOrderIdKey := fmt.Sprintf("%s%v", cacheTradeOrderIdPrefix, data.Id)
	tradeOrderTradeOrderIdKey := fmt.Sprintf("%s%v", cacheTradeOrderTradeOrderIdPrefix, data.TradeOrderId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tradeOrderRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.TradeOrderId, newData.PayOrderId, newData.CarbonAssetId, newData.Initiator, newData.Recipient, newData.TadeStatus, newData.InitiatorTime, newData.FinishTime, newData.Id)
	}, tradeOrderIdKey, tradeOrderTradeOrderIdKey)
	return err
}

func (m *defaultTradeOrderModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTradeOrderIdPrefix, primary)
}

func (m *defaultTradeOrderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tradeOrderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTradeOrderModel) tableName() string {
	return m.table
}