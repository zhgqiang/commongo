package db

import (
	"fmt"

	"google.golang.org/grpc"
)

// DGraph is 数据库配置.
type DGraph struct {
	Host string `json:"host" toml:"host" description:"数据库地址"`
	Port int    `json:"port" toml:"port" description:"数据库端口"`
}

// NewConn is 创建数据库连接.
func (p *DGraph) NewConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", p.Host, p.Port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
