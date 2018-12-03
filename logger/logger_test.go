package logger_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/zhgqiang/common/data"
	"github.com/zhgqiang/common/logger"
)

var defaultConfig = `
{
	"level": "debug",
	"writerMap": {
		"debug": "console",
		"info": "console",
		"warn": "channel",
		"error": "channel"
	},
	"file": {
		"path": "logs",
		"maxAgeHour": 720,
		"rotationTimeHour": 24
	}
}
`

func TestNewChannelLogger(t *testing.T) {
	c := make(chan data.JSON)
	// done := make(chan struct{})
	go func() {
		for {
			select {
			// case <-done:
			// 	return
			case b, ok := <-c:
				if !ok {
					fmt.Println("数据错误")
					return
				}
				fmt.Println("输出数据", string(b))
			}
		}
	}()
	config := logger.Logger{}
	err := json.Unmarshal([]byte(defaultConfig), &config)
	if err != nil {
		t.Fatal("配置信息解析错误", err)
	}
	err = logger.NewChannelLogger(config, c)
	if err != nil {
		t.Fatal("配置信息错误", err)
	}

	logrus.Debug(1)
	logrus.Info(2)
	logrus.Warn(3)
	logrus.Error(4)
	time.Sleep(time.Second * 10)
}
