package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"hub/internal/config"
	"hub/internal/middleware"
	"hub/model"
	"log"
	"pay/payserver"
	"retire_cert"
	"trade/tradeserver"
)

type ServiceContext struct {
	Config              config.Config
	Redis               *redis.Redis
	ServiceRpc          ServiceRpc
	Middleware          Middleware
	RetireCert          RetireCert
	MysqlServiceContext *MysqlServiceContext
}

type RetireCert struct {
	RetireCertInter retire_cert.RetireConfigInter
}

type ServiceRpc struct {
	//GsfRpc   gsfserver.GsfServer
	PayRpc   payserver.PayServer
	TradeRpc tradeserver.TradeServer
}

type MysqlServiceContext struct {
	UserDB     sqlx.SqlConn
	Assets     model.AssetsModel
	AssetsSell model.AssetsSellModel
	Retire     model.RetireModel
	UserWallet model.UserWalletModel
}

type Middleware struct {
	AuthJwtMiddleware   rest.Middleware
	WalletJwtMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlService := MysqlService(c)
	newRedis, err := redis.NewRedis(c.CacheRedis[0].RedisConf)
	if err != nil {
		log.Println(err)
		return nil
	}

	client, err := zrpc.NewClient(c.ServiceRpc.PayRpc)
	if err != nil {
		//return nil
	}

	client2, err := zrpc.NewClient(c.ServiceRpc.TradeRpc)
	if err != nil {
		//return nil
	}

	return &ServiceContext{
		Config: c,
		RetireCert: RetireCert{
			RetireCertInter: retire_cert.NewRetireC(c.RetireCert.UploadPath, c.RetireCert.ImageBackground, c.RetireCert.FontTTF),
		},
		ServiceRpc: ServiceRpc{
			//GsfRpc:   gsfserver.NewGsfServer(zrpc.MustNewClient(c.ServiceRpc.GsfRpc)),
			PayRpc:   payserver.NewPayServer(client),
			TradeRpc: tradeserver.NewTradeServer(client2),
		},
		Middleware: Middleware{
			AuthJwtMiddleware:   middleware.NewAuthJwtMiddleware(c.ServiceJwtSign.UserServiceAuth, mysqlService.UserDB).Handle,
			WalletJwtMiddleware: middleware.NewWalletJwtMiddleware(c.ServiceJwtSign.HubServiceAuth, mysqlService.UserWallet).Handle,
		},
		MysqlServiceContext: mysqlService,
		Redis:               newRedis,
	}
}

func MysqlService(c config.Config) *MysqlServiceContext {
	var (
		userDB, hubDb sqlx.SqlConn
		err           error
	)
	userDB = sqlx.NewMysql(c.MysqlService.UserService.DataSourceName)
	if err != nil {
		panic(err)
	}

	hubDb = sqlx.NewMysql(c.MysqlService.HubService.DataSourceName)
	if err != nil {
		panic(err)
	}

	//tradeDb = sqlx.NewMysql(c.MysqlService.TradeMysql.DataSourceName)
	//if err != nil {
	//	panic(err)
	//}
	//
	//payDb = sqlx.NewMysql(c.MysqlService.PayMysql.DataSourceName)
	//if err != nil {
	//	panic(err)
	//}

	return &MysqlServiceContext{
		UserDB:     userDB,
		Assets:     model.NewAssetsModel(hubDb),
		AssetsSell: model.NewAssetsSellModel(hubDb, c.CacheRedis),
		Retire:     model.NewRetireModel(hubDb),
		UserWallet: model.NewUserWalletModel(hubDb),
	}
}
