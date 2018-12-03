package utils

import (
	"encoding/json"
	"errors"
)

// Convert 将 data 数据 转为 result 数据
func Convert(result interface{}, data interface{}) error {
	if data == nil {
		return errors.New("数据为空")
	}
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, result)
	if err != nil {
		return err
	}
	return nil
}
