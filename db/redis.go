package db

import (
	"net"
	"strconv"

	"github.com/go-redis/redis"
)

// Redis is 内存数据库配置
type Redis struct {
	Host        string `json:"host" toml:"host" description:"数据库地址"`
	Port        int    `json:"port" toml:"port" description:"数据库端口"`
	Password    string `json:"password" toml:"password" description:"数据库访问密码"`
	DB          int    `json:"db" toml:"db" description:"启用的数据库"`
	PoolSize    int    `json:"poolSize" toml:"poolSize" description:"连接池大小"`
	IdleTimeout int64  `json:"idleTimeout" toml:"idleTimeout" description:"最大空闲超时时间"`
}

// NewClient is 创建内存数据库客户端
func (p *Redis) NewClient() *redis.Client {
	c := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(p.Host, strconv.Itoa(p.Port)),
		Password: p.Password, // no password set
		DB:       p.DB,       // use default DB
		PoolSize: p.PoolSize,
	})
	return c
}
