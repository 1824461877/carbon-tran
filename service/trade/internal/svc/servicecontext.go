package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"trade/internal/config"
	"trade/model"
)

type ServiceContext struct {
	Config     config.Config
	UserDB     sqlx.SqlConn
	TradeOrder model.TradeOrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		sqlDB, userDB sqlx.SqlConn
		err           error
	)

	userDB = sqlx.NewMysql(c.MysqlService.UserMysql.DataSourceName)
	if err != nil {
		panic(err)
	}

	sqlDB = sqlx.NewMysql(c.MysqlService.TradeMysql.DataSourceName)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:     c,
		UserDB:     userDB,
		TradeOrder: model.NewTradeOrderModel(sqlDB, c.CacheRedis),
	}
}
