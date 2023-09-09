package svc

import (
	"github.com/plutov/paypal"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	paypal2 "pay/common/paypal"
	"pay/internal/config"
	"pay/model"
	"trade/tradeserver"
)

type ServiceContext struct {
	Config        config.Config
	Redis         *redis.Redis
	PayPalService *paypal2.PayClient
	PayOrder      model.PayOrderModel
	UserWallet    model.UserWalletModel
	ServiceRpc    ServiceRpc
}

type ServiceRpc struct {
	TradeRpc tradeserver.TradeServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	var (
		walletDB, payDB sqlx.SqlConn
		err             error
	)

	//walletDB = sqlx.NewMysql(c.MysqlService.WalletMysql.DataSourceName)
	//if err != nil {
	//	panic(err)
	//}

	newRedis, err := redis.NewRedis(c.CacheRedis[0].RedisConf)
	if err != nil {
		log.Println(err)
		return nil
	}

	payDB = sqlx.NewMysql(c.MysqlService.PayMysql.DataSourceName)
	if err != nil {
		panic(err)
	}

	walletDB = sqlx.NewMysql(c.MysqlService.WalletMysql.DataSourceName)
	if err != nil {
		panic(err)
	}

	var (
		payClient *paypal2.PayClient
	)
	switch c.ServiceMode.PayPalMode.Open {
	case "Sandbox":
		payClient = paypal2.NewPayClient(&paypal2.PayClientConfig{
			ClientID: c.ServiceMode.PayPalMode.Sandbox.ClientID,
			Secret:   c.ServiceMode.PayPalMode.Sandbox.Secret,
			APIBase:  paypal.APIBaseSandBox,
		})
	case "Live":
		payClient = paypal2.NewPayClient(&paypal2.PayClientConfig{
			ClientID: c.ServiceMode.PayPalMode.Live.ClientID,
			Secret:   c.ServiceMode.PayPalMode.Live.Secret,
			APIBase:  paypal.APIBaseLive,
		})
	}

	client, err := zrpc.NewClient(c.ServiceRpc.TradeRpc)
	if err != nil {
		//return nil
	}

	return &ServiceContext{
		Config:        c,
		PayPalService: payClient,
		ServiceRpc: ServiceRpc{
			TradeRpc: tradeserver.NewTradeServer(client),
		},
		UserWallet: model.NewUserWalletModel(walletDB),
		PayOrder:   model.NewPayOrderModel(payDB, c.CacheRedis),
		Redis:      newRedis,
	}
}
