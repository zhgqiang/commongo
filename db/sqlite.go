package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// SQLite is 数据库配置.
type SQLite struct {
	DriverName     string `json:"driverName" toml:"driverName" description:"驱动名称"`
	DataSourceName string `json:"dataSourceName" toml:"dataSourceName" description:"数据源名称"`
	MaxOpenConns   int    `json:"maxOpenConns" toml:"maxOpenConns" description:"数据库最大连接数"`
	MaxIdleConns   int    `json:"maxIdleConns" toml:"maxIdleConns" description:"数据库最大空闲连接数"`
	IdleTime       int64  `json:"idleTime" toml:"idleTime" description:"数据库最大空闲时间"`

	db *gorm.DB
}

// NewConn is 创建数据库连接.
func (p *SQLite) NewConn() (*gorm.DB, error) {
	db, err := gorm.Open(p.DriverName, p.DataSourceName)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(p.MaxIdleConns)
	db.DB().SetMaxOpenConns(p.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(p.IdleTime) * time.Second)
	return db, nil
}

// Init is 初始化数据库连接.
func (p *SQLite) Init() error {
	db, err := gorm.Open(p.DriverName, p.DataSourceName)
	if err != nil {
		return err
	}
	db.DB().SetMaxIdleConns(p.MaxIdleConns)
	db.DB().SetMaxOpenConns(p.MaxOpenConns)
	p.db = db
	return nil
}

// GetSQL is 通过 sql 语句查询数据库.
func (p *SQLite) GetSQL(sql string) ([]map[string]interface{}, error) {
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
