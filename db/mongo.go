package db

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// MongoDB is 数据库配置.
type MongoDB struct {
	Host     string `json:"host" toml:"host" description:"数据库地址"`
	Port     int    `json:"port" toml:"port" description:"数据库端口"`
	Username string `json:"username" toml:"username" description:"数据库访问用户名"`
	Password string `json:"password" toml:"password" description:"数据库访问密码"`
	Database string `json:"database" toml:"database" description:"数据库名称"`
}

// NewConn is 创建数据库连接.
func (p *MongoDB) NewConn() (*mongo.Client, error) {
	client, err := mongo.NewClient(fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", p.Username, p.Password, p.Host, p.Port, p.Database))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
