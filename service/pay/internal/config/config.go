package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CacheRedis     cache.CacheConf
	PayServiceAuth PayServiceAuth
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
