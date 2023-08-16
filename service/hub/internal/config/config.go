package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CacheRedis     cache.CacheConf
	ServiceJwtSign ServiceJwtSign
	MysqlService   MysqlService
	ServiceRpc     ServiceRpc
}

type ServiceRpc struct {
	GsfRpc   zrpc.RpcClientConf
	PayRpc   zrpc.RpcClientConf
	TradeRpc zrpc.RpcClientConf
}

type ServiceJwtSign struct {
	UserServiceAuth  ServiceAuth
	HubServiceAuth   ServiceAuth
	PayServiceAuth   ServiceAuth
	TradeServiceAuth ServiceAuth
}

type ServiceAuth struct {
	JwtSignKey    string
	JwtSignExpire int64
}

type MysqlService struct {
	UserService MysqlSource
	HubService  MysqlSource
}

type MysqlSource struct {
	DataSourceName string
}
