package logger

import (
	"errors"

	"github.com/zhgqiang/commongo/data"
)

// ChannelWriter is 配置日志输出位置为管道中
type ChannelWriter struct {
	c chan data.JSON
}

// NewChanWriter is 创建管道输出
func NewChanWriter(channel chan data.JSON) (*ChannelWriter, error) {
	return &ChannelWriter{c: channel}, nil
}

// ChannelWriter Write is 向管道中写入数据
func (p *ChannelWriter) Write(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, errors.New("data length is 0")
	}
	p.c <- b
	return len(b), err
}
