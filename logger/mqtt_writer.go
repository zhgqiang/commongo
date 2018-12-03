package logger

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"

	"github.com/zhgqiang/common/mq"
)

// EmqttWriter is 配置日志输出位置为消息队列
type EmqttWriter struct {
	ops   mq.Emqtt
	topic string
}

// NewEmqttWriter is 创建消息队列输出
func NewEmqttWriter(ops mq.Emqtt, topic string) (*EmqttWriter, error) {
	return &EmqttWriter{ops: ops, topic: topic}, nil
}

// EmqttWriter Write is 向消息队列中写入数据
func (rl *EmqttWriter) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, errors.New("数据为空")
	}
	conn, err := net.Dial("tcp", net.JoinHostPort(rl.ops.Host, strconv.Itoa(rl.ops.Port)))
	if err != nil {
		return 0, err
	}
	cc := mqtt.NewClientConn(conn)
	if err := cc.Connect(rl.ops.Username, rl.ops.Password); err != nil {
		return 0, fmt.Errorf("EMQTT连接错误.%+v", err)
	}

	cc.Publish(&proto.Publish{
		TopicName: rl.topic,
		Payload:   proto.BytesPayload(p),
	})
	cc.Disconnect()
	return len(p), err
}
