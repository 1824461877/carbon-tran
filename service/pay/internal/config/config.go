package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CacheRedis     cache.CacheConf
	PayServiceAuth PayServiceAuth
	ServiceMode    ServiceMode
	ServiceRpc     ServiceRpc
	MysqlService   MysqlService
}

type MysqlService struct {
	WalletMysql Mysql
	PayMysql    Mysql
}

type Mysql struct {
	DataSourceName string
}

type PayServiceAuth struct {
	UserJwtSignKey    string
	UserJwtSignExpire int64
	JwtSignKey        string
	JwtSignExpire     int64
}

type ServiceMode struct {
	PayPalMode PayPalMode
}

type PayPalMode struct {
	Open    string
	Sandbox PalMode
	Live    PalMode
}

type PalMode struct {
	ClientID string
	Secret   string
}

type ServiceRpc struct {
	TradeRpc zrpc.RpcClientConf
}
