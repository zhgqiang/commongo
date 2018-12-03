package db

import (
	"database/sql"
	"fmt"
	"time"
)

// Oracle is 数据库配置.
type Oracle struct {
	Host         string `json:"host" toml:"host" description:"数据库地址"`
	Port         int    `json:"port" toml:"port" description:"数据库端口"`
	Username     string `json:"username" toml:"username" description:"数据库访问用户名"`
	Password     string `json:"password" toml:"password" description:"数据库访问密码"`
	SID          string `json:"sid" toml:"sid" description:"数据库实例名称"`
	MaxOpenConns int    `json:"maxOpenConns" toml:"maxOpenConns" description:"数据库最大连接数"`
	MaxIdleConns int    `json:"maxIdleConns" toml:"maxIdleConns" description:"数据库最大空闲连接数"`
	IdleTime     int64  `json:"idleTime" toml:"idleTime" description:"数据库最大空闲时间"`

	db *sql.DB
}

// NewConn is 创建数据库连接.
func (p *Oracle) NewConn() (*sql.DB, error) {
	db, err := sql.Open("oci8", fmt.Sprintf("%s/%s@%s:%d/%s", p.Username, p.Password, p.Host, p.Port, p.SID))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(p.MaxOpenConns)
	db.SetMaxIdleConns(p.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(p.IdleTime) * time.Second)
	return db, nil
}

// Init is 初始化数据库连接.
func (p *Oracle) Init() error {
	db, err := sql.Open("oci8", fmt.Sprintf("%s/%s@%s:%d/%s", p.Username, p.Password, p.Host, p.Port, p.SID))
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(p.MaxOpenConns)
	db.SetMaxIdleConns(p.MaxIdleConns)
	p.db = db
	return nil
}

// GetSQL is 通过 sql 语句查询数据库.
func (p *Oracle) GetSQL(sql string) ([]map[string]interface{}, error) {
	rows, err := p.db.Query(sql)
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
