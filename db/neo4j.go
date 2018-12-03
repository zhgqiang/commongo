package db

import (
	"fmt"

	"github.com/jmcvetta/neoism"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Neo4j is 图形数据库配置
type Neo4j struct {
	Host        string `json:"host" toml:"host" description:"数据库地址"`
	Port        int    `json:"port" toml:"port" description:"数据库端口"`
	Username    string `json:"username" toml:"username" description:"数据库访问用户名"`
	Password    string `json:"password" toml:"password" description:"数据库访问密码"`
	BoltPort    int    `json:"boltPort" toml:"boltPort" descriptioFn:"bolt端口"`
	BoltPoolMax int    `json:"boltPoolMax" toml:"boltPoolMax" descriptioFn:"bolt连接池连接数量"`
}

// NewConn is 创建普通连接
func (p *Neo4j) NewConn() (db *neoism.Database, err error) {
	db, err = neoism.Connect(fmt.Sprintf("http://%s:%s@%s:%d/db/data", p.Username, p.Password, p.Host, p.Port))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// NewBoltConn is 创建 Bolt 连接
func (p *Neo4j) NewBoltConn() (driverPool golangNeo4jBoltDriver.DriverPool, err error) {
	driverPool, err = golangNeo4jBoltDriver.NewDriverPool(fmt.Sprintf("bolt://%s:%s@%s:%d", p.Username, p.Password, p.Host, p.BoltPort), p.BoltPoolMax)
	if err != nil {
		return nil, err
	}
	return driverPool, nil
}
