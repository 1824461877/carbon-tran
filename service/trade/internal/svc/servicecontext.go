package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"pay/payserver"
	"trade/internal/config"
	"trade/model"
)

type ServiceContext struct {
	Config              config.Config
	ServiceRpc          ServiceRpc
	MysqlServiceContext MysqlServiceContext
	Hub                 model.TradeOrderModel
}

type MysqlServiceContext struct {
	UserDB     sqlx.SqlConn
	TradeOrder model.TradeOrderModel
	Assets     model.AssetsModel
	AssetsSell model.AssetsSellModel
}

type ServiceRpc struct {
	PayRpc payserver.PayServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		hubDB, tradeDB, userDB sqlx.SqlConn
		err                    error
	)

	userDB = sqlx.NewMysql(c.MysqlService.UserMysql.DataSourceName)
	if err != nil {
		panic(err)
	}

	tradeDB = sqlx.NewMysql(c.MysqlService.TradeMysql.DataSourceName)
	if err != nil {
		panic(err)
	}

	hubDB = sqlx.NewMysql(c.MysqlService.HubMysql.DataSourceName)
	if err != nil {
		panic(err)
	}

	client, err := zrpc.NewClient(c.ServiceRpc.PayRpc)
	if err != nil {
		//panic(err)
	}

	return &ServiceContext{
		Config: c,
		MysqlServiceContext: MysqlServiceContext{
			UserDB:     userDB,
			TradeOrder: model.NewTradeOrderModel(tradeDB),
			Assets:     model.NewAssetsModel(hubDB),
			AssetsSell: model.NewAssetsSellModel(hubDB, c.CacheRedis),
		},
		ServiceRpc: ServiceRpc{
			PayRpc: payserver.NewPayServer(client),
		},
	}
}
