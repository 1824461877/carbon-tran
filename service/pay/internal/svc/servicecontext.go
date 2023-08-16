package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"pay/internal/config"
	"pay/model"
)

type ServiceContext struct {
	Config     config.Config
	UserWallet model.UserWalletModel
	PayOrder   model.PayOrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		walletDB, payDB sqlx.SqlConn
		err             error
	)

	walletDB = sqlx.NewMysql(c.MysqlService.WalletMysql.DataSourceName)
	if err != nil {
		panic(err)
	}

	payDB = sqlx.NewMysql(c.MysqlService.PayMysql.DataSourceName)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:     c,
		UserWallet: model.NewUserWalletModel(walletDB),
		PayOrder:   model.NewPayOrderModel(payDB, c.CacheRedis),
	}
}
