// Package mq 消息队列操作.
package mq

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"
)

// Emqtt is 配置信息.
type Emqtt struct {
	Host      string `json:"host" toml:"host" description:"EMQTT地址"`
	Port      int    `json:"port" toml:"port" description:"EMQTT端口"`
	Username  string `json:"username" toml:"username" description:"EMQTT访问用户名"`
	Password  string `json:"password" toml:"password" description:"EMQTT访问密码"`
	TopicName string `json:"topicName" toml:"topicName" description:"EMQTT TopicName"`
}

// SendKeyValue is Emqtt 发送key、value.
func (p *Emqtt) SendKeyValue(key string, val interface{}) (err error) {
	conn, err := net.Dial("tcp", net.JoinHostPort(p.Host, strconv.Itoa(p.Port)))
	if err != nil {
		return err
	}
	cc := mqtt.NewClientConn(conn)
	if err := cc.Connect(p.Username, p.Password); err != nil {
		return fmt.Errorf("EMQTT连接错误.%v", err)
	}
	m := map[string]interface{}{key: val}
	mb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	cc.Publish(&proto.Publish{
		TopicName: p.TopicName,
		Payload:   proto.BytesPayload(mb),
	})
	cc.Disconnect()
	return nil
}

// SendTopicKeyValue is Emqtt 向 topic 发送key、value.
func (p *Emqtt) SendTopicKeyValue(topic, key string, val interface{}) (err error) {
	conn, err := net.Dial("tcp", net.JoinHostPort(p.Host, strconv.Itoa(p.Port)))
	if err != nil {
		return err
	}
	cc := mqtt.NewClientConn(conn)
	if err := cc.Connect(p.Username, p.Password); err != nil {
		return fmt.Errorf("EMQTT连接错误.%v", err)
	}
	m := map[string]interface{}{key: val}
	mb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	cc.Publish(&proto.Publish{
		TopicName: p.TopicName,
		Payload:   proto.BytesPayload(mb),
	})
	cc.Disconnect()
	return nil
}

// SendTopicValue is Emqtt 发送topic、value.
func (p *Emqtt) SendTopicValue(topic string, val interface{}) (err error) {
	conn, err := net.Dial("tcp", net.JoinHostPort(p.Host, strconv.Itoa(p.Port)))
	if err != nil {
		return err
	}
	cc := mqtt.NewClientConn(conn)
	if err := cc.Connect(p.Username, p.Password); err != nil {
		return fmt.Errorf("EMQTT连接错误.%v", err)
	}
	mb, err := json.Marshal(val)
	if err != nil {
		return err
	}
	cc.Publish(&proto.Publish{
		TopicName: topic,
		Payload:   proto.BytesPayload(mb),
	})
	cc.Disconnect()
	return nil
}

// Send is Emqtt 发送 msg 数据.
func (p *Emqtt) Send(msg string) (err error) {
	conn, err := net.Dial("tcp", net.JoinHostPort(p.Host, strconv.Itoa(p.Port)))
	if err != nil {
		return err
	}
	cc := mqtt.NewClientConn(conn)
	if err := cc.Connect(p.Username, p.Password); err != nil {
		return fmt.Errorf("EMQTT连接错误.%v", err)
	}
	cc.Publish(&proto.Publish{
		TopicName: p.TopicName,
		Payload:   proto.BytesPayload([]byte(msg)),
	})
	cc.Disconnect()
	return nil
}
