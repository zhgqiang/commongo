package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

var sample = `
{
	"host": "101.200.39.236",
	"port": 6379,
	"username": "root",
	"password": "dell@123",
	"database": "cpgd",
	"maxOpenConns": 200,
	"maxIdleConns": 100,
}
`

// Mariadb is 数据库配置.
type Mariadb struct {
	Host         string `json:"host" toml:"host" description:"数据库地址"`
	Port         int    `json:"port" toml:"port" description:"数据库端口"`
	Username     string `json:"username" toml:"username" description:"数据库访问用户名"`
	Password     string `json:"password" toml:"password" description:"数据库访问密码"`
	Database     string `json:"database" toml:"database" description:"数据库名称"`
	MaxOpenConns int    `json:"maxOpenConns" toml:"maxOpenConns" description:"数据库最大连接数"`
	MaxIdleConns int    `json:"maxIdleConns" toml:"maxIdleConns" description:"数据库最大空闲连接数"`
	IdleTime     int64  `json:"idleTime" toml:"idleTime" description:"数据库最大空闲时间"`

	db *gorm.DB
}

// NewConn is 创建数据库连接.
func (p *Mariadb) NewConn() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", p.Username, p.Password, p.Host, p.Port, p.Database))
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(p.MaxIdleConns)
	db.DB().SetMaxOpenConns(p.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(p.IdleTime) * time.Second)
	return db, nil
}

// Init is 初始化数据库连接.
func (p *Mariadb) Init() error {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", p.Username, p.Password, p.Host, p.Port, p.Database))
	if err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(p.MaxIdleConns)
	db.DB().SetMaxOpenConns(p.MaxOpenConns)
	p.db = db
	return nil
}

// GetSQL is 通过 sql 语句查询数据库.
func (p *Mariadb) GetSQL(sql string) ([]map[string]interface{}, error) {
	rows, err := p.db.Raw(sql).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	size := len(columns)
	pts := make([]interface{}, size)
	c := make([]interface{}, size)
	container := make([]map[string]interface{}, 0)
	for i := range pts {
		pts[i] = &c[i]
	}
	for rows.Next() {
		err = rows.Scan(pts...)
		if err != nil {
			return nil, err
		}
		var r = make(map[string]interface{}, size)
		for i, column := range columns {
			val := pts[i].(*interface{})
			b, ok := (*val).([]byte)
			if ok {
				v := string(b)
				r[column] = v
			} else {
				r[column] = *val
			}
		}
		container = append(container, r)
	}
	return container, nil
}
