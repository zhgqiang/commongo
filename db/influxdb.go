package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/influxdata/influxdb/client/v2"

	"github.com/zhgqiang/commongo/data"
)

// InfluxDB is 数据库配置.
type InfluxDB struct {
	Addr     string `json:"addr" toml:"addr" description:"数据库地址"`
	Username string `json:"username" toml:"username" description:"数据库访问用户名"`
	Password string `json:"password" toml:"password" description:"数据库访问密码"`
	Database string `json:"database" toml:"database" description:"数据库名"`

	client client.Client
}

// NewClient is 创建数据库客户端.
func (p *InfluxDB) NewClient() (c client.Client, err error) {
	// Create a new HTTPClient
	return client.NewHTTPClient(client.HTTPConfig{
		Addr:     p.Addr,
		Username: p.Username,
		Password: p.Password,
	})
}

// Init is 初始化数据库连接.
func (p *InfluxDB) Init() error {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     p.Addr,
		Username: p.Username,
		Password: p.Password,
	})
	if err != nil {
		return err
	}
	p.client = cli
	return nil
}

// CreateDB is 创建数据库
func (p *InfluxDB) CreateDB(database ...string) (err error) {
	if database != nil && len(database) > 0 {
		for _, db := range database {
			_, err = p.QueryDB(fmt.Sprintf("CREATE DATABASE %s", db))
			if err != nil {
				return fmt.Errorf("创建数据库失败.%v", err)
			}
		}
	} else {
		_, err = p.QueryDB(fmt.Sprintf("CREATE DATABASE %s", p.Database))
		if err != nil {
			return fmt.Errorf("创建数据库失败.%v", err)
		}
	}
	return nil
}

// QueryDB convenience function to query the database
func (p *InfluxDB) QueryDB(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: p.Database,
	}

	if response, err := p.client.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// Get is 查询返回基本数据
func (p *InfluxDB) Get(res []client.Result, timeFormat string) (*[][]data.Base, error) {
	if res == nil || len(res) == 0 || len(res[0].Series) == 0 || len(res[0].Series[0].Values) == 0 {
		return nil, errors.New("查询结果为空")
	}
	arr := make([][]data.Base, len(res[0].Series[0].Values[0])-1)
	for i, row := range res[0].Series[0].Values {
		for j := range res[0].Series[0].Columns {
			if j == 0 {
				t, err := time.Parse(time.RFC3339, row[0].(string))
				if err != nil {
					return nil, errors.New("查询出的时间转time.Time错误")
				}
				for k := 0; k < len(row)-1; k++ {
					arr[k] = append(arr[k], data.Base{Name: t.Local().Format(timeFormat)})
				}
			} else {
				arr[j-1][i].Value = row[j]
			}
		}
	}
	return &arr, nil
}

// CreateRetentionPolicy is 创建保留策略
func (p *InfluxDB) CreateRetentionPolicy(policyName, duration string, replication int) error {
	_, err := p.QueryDB(fmt.Sprintf("CREATE RETENTION POLICY \"%s\" ON %s DURATION %s REPLICATION %d", policyName, p.Database, duration, replication))
	if err != nil {
		return fmt.Errorf("创建保留策略失败.%v", err)
	}
	return nil
}

// WritePointsRetention is 创建保留策略
func (p *InfluxDB) WritePointsRetention(precision, policyName, measurement string, tags map[string]string, fields map[string]interface{}, timestamp time.Time) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:        p.Database,
		Precision:       precision,
		RetentionPolicy: policyName,
	})
	if err != nil {
		return fmt.Errorf("新建批量添加操作失败(Retention).%v", err)
	}

	pt, err := client.NewPoint(
		measurement,
		tags,
		fields,
		timestamp,
	)
	if err != nil {
		return fmt.Errorf("新建Point失败(Retention).%v", err)
	}
	bp.AddPoint(pt)

	if err := p.client.Write(bp); err != nil {
		return fmt.Errorf("存储Point失败(Retention).%v", err)
	}
	return nil
}

// WritePointsRetentionBatch is 新建批量添加操作
func (p *InfluxDB) WritePointsRetentionBatch(precision, policyName, measurement string, tags []map[string]string, fields []map[string]interface{}, timestamps []time.Time) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:        p.Database,
		Precision:       precision,
		RetentionPolicy: policyName,
	})
	if err != nil {
		return fmt.Errorf("新建批量添加操作失败(Retention).%v", err)
	}

	for i := 0; i < len(tags); i++ {
		pt, err := client.NewPoint(
			measurement,
			tags[i],
			fields[i],
			timestamps[i],
		)
		if err != nil {
			return fmt.Errorf("新建Point失败(Retention).%v", err)
		}
		bp.AddPoint(pt)
	}

	if err := p.client.Write(bp); err != nil {
		return fmt.Errorf("存储Point失败(Retention).%v", err)
	}
	return nil
}
