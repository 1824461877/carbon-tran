package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CacheRedis       cache.CacheConf
	TradeServiceAuth TradeServiceAuth
	MysqlService     MysqlService
}

type MysqlService struct {
	UserMysql  Mysql
	TradeMysql Mysql
}

type Mysql struct {
	DataSourceName string
}

type TradeServiceAuth struct {
	UserJwtSignKey    string
	UserJwtSignExpire int64
	JwtSignKey        string
	JwtSignExpire     int64
}
