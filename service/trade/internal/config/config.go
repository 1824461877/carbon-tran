package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CacheRedis     cache.CacheConf
	ServiceJwtSign ServiceJwtSign
	ServiceRpc     ServiceRpc
	MysqlService   MysqlService
}

type ServiceRpc struct {
	PayRpc zrpc.RpcClientConf
}

type MysqlService struct {
	HubMysql   Mysql
	UserMysql  Mysql
	TradeMysql Mysql
}

type Mysql struct {
	DataSourceName string
}

type ServiceJwtSign struct {
	UserServiceAuth  ServiceAuth
	PayServiceAuth   ServiceAuth
	TradeServiceAuth ServiceAuth
}

type ServiceAuth struct {
	JwtSignKey    string
	JwtSignExpire int64
}
