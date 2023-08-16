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
	assetsFieldNames          = builder.RawFieldNames(&Assets{})
	assetsRows                = strings.Join(assetsFieldNames, ",")
	assetsRowsExpectAutoSet   = strings.Join(stringx.Remove(assetsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	assetsRowsWithPlaceHolder = strings.Join(stringx.Remove(assetsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	assetsModel interface {
		Insert(ctx context.Context, data *Assets) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Assets, error)
		FindUidAll(ctx context.Context, uid string) (*[]Assets, error)
		FindUidOne(ctx context.Context, uid string) (*Assets, error)
		FindAssIdOne(ctx context.Context, assId string) (*Assets, error)
		Update(ctx context.Context, data *Assets) error
		UpdateListing(ctx context.Context, data *Assets) error
		UpdateNumber(ctx context.Context, data *Assets) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAssetsModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Assets struct {
		Id           int64     `db:"id"`
		AssId        string    `db:"ass_id"`        // 用于api操作
		Hid          string    `db:"hid"`           // 资产唯一标识
		UserId       string    `db:"user_id"`       // 资产所属用户
		Source       string    `db:"source"`        // 来源资产
		Number       int64     `db:"number"`        // 数量
		Project      string    `db:"project"`       // 项目名称
		GsId         string    `db:"gs_id"`         // gs_id
		RetireNumber int64     `db:"retire_number"` // 注销数量
		Status       string    `db:"status"`        // status
		Country      string    `db:"country"`       // country
		ProjectType  string    `db:"project_type"`  // country
		Product      string    `db:"product"`       // product
		VersHead     int64     `db:"vers_head"`     // vers 序号区间 （head）
		VersTail     int64     `db:"vers_tail"`     // vers 序号区间 （tail）
		SerialNumber string    `db:"serial_number"` // 资产编号
		Day          int64     `db:"day"`           // 日期时间
		Listing      bool      `db:"listing"`       // 挂牌
		CreateTime   time.Time `db:"create_time"`   // 资产创建时间
	}
)

func newAssetsModel(conn sqlx.SqlConn) *defaultAssetsModel {
	return &defaultAssetsModel{
		conn:  conn,
		table: "`assets`",
	}
}

func (m *defaultAssetsModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAssetsModel) FindOne(ctx context.Context, id int64) (*Assets, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", assetsRows, m.table)
	var resp Assets
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

func (m *defaultAssetsModel) Insert(ctx context.Context, data *Assets) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, assetsRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.AssId, data.Hid, data.UserId, data.Source, data.Number, data.Project, data.GsId, data.Status, data.Country, data.ProjectType, data.Product, data.VersHead, data.VersTail, data.SerialNumber, data.Day, data.Listing)
	return ret, err
}

func (m *defaultAssetsModel) FindUidOne(ctx context.Context, uid string) (*Assets, error) {
	query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", assetsRows, m.table)
	var resp Assets
	err := m.conn.QueryRowCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAssetsModel) FindUidAll(ctx context.Context, uid string) (*[]Assets, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ?", assetsRows, m.table)
	var resp []Assets
	err := m.conn.QueryRowsCtx(ctx, &resp, query, uid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAssetsModel) FindAssIdOne(ctx context.Context, assId string) (*Assets, error) {
	query := fmt.Sprintf("select %s from %s where `ass_id` = ? limit 1", assetsRows, m.table)
	var resp Assets
	err := m.conn.QueryRowCtx(ctx, &resp, query, assId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAssetsModel) Update(ctx context.Context, data *Assets) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, assetsRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.AssId, data.Hid, data.UserId, data.Source, data.Number, data.Project, data.GsId, data.Status, data.Country, data.ProjectType, data.Product, data.VersHead, data.VersTail, data.SerialNumber, data.Day,data.Listing, data.Id)
	return err
}

func (m *defaultAssetsModel) UpdateListing(ctx context.Context, data *Assets) error {
	query := fmt.Sprintf("update %s set `listing`= ? where `ass_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query,data.Listing, data.AssId)
	return err
}

func (m *defaultAssetsModel) UpdateNumber(ctx context.Context, data *Assets) error {
	query := fmt.Sprintf("update %s set `number`= ?,`retire_number`= ? where `ass_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query,data.Number,data.RetireNumber,data.AssId)
	return err
}

func (m *defaultAssetsModel) tableName() string {
	return m.table
}
