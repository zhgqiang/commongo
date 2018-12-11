// Package logger holds 日志配置及处理
package logger

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"github.com/zhgqiang/commongo/data"
	"github.com/zhgqiang/commongo/mq"
)

const (
	// CONSOLE is 控制台输出.
	CONSOLE LogLoc = "console"
	// FILE is 文件输出.
	FILE LogLoc = "file"
	// CHANNEL is 将日志输出到一个 chan []byte 管道.
	CHANNEL LogLoc = "channel"
	// EMQTT is 将日志输出到消息队列.
	EMQTT LogLoc = "emqtt"

	// DEBUG is  级别最低，可以随意的使用于任何觉得有利于在调试时更详细的了解系统运行状态的东东.
	DEBUG LogLevel = "debug"
	// INFO is 重要，输出信息,用来反馈系统的当前状态给最终用户的.
	INFO LogLevel = "info"
	// WARN is 可修复，系统可继续运行下去.
	WARN LogLevel = "warn"
	// ERROR is 可修复性，但无法确定系统会正常的工作下去.
	ERROR LogLevel = "error"
)

// LogLoc is 日志输出位置.
type LogLoc string

// LogLevel is 日志输出等级.
type LogLevel string

// Logger is 日志配置.
type Logger struct {
	Level     LogLevel `json:"level" toml:"level"`
	WriterMap struct {
		Debug LogLoc `json:"debug" toml:"debug"`
		Info  LogLoc `json:"info" toml:"info"`
		Warn  LogLoc `json:"warn" toml:"warn"`
		Error LogLoc `json:"error" toml:"error"`
	} `json:"writerMap" toml:"writerMap"`
	File struct {
		Path             string `json:"path" toml:"path"`
		MaxAgeHour       int    `json:"maxAgeHour" toml:"maxAgeHour"`
		RotationTimeHour int    `json:"rotationTimeHour" toml:"rotationTimeHour"`
	} `json:"file" toml:"file"`
}

// NewLogger is 初始化普通日志配置等.
func NewLogger(config Logger) error {
	writerMap := lfshook.WriterMap{}
	var fileWriter io.Writer
	if &config.File != nil {
		_, statErr := os.Stat(config.File.Path)
		// 如果目录不存在则创建该目录
		if os.IsNotExist(statErr) {
			if err := os.MkdirAll(config.File.Path, 0755); err != nil {
				return err
			}
		}
		writer, err := rotatelogs.New(
			path.Join(config.File.Path, "%Y%m%d%H%M.log"),
			rotatelogs.WithMaxAge(time.Hour*time.Duration(config.File.MaxAgeHour)),             // 文件最大保存时间
			rotatelogs.WithRotationTime(time.Hour*time.Duration(config.File.RotationTimeHour)), // 日志切割时间间隔
		)
		if err != nil {
			return err
		}
		fileWriter = writer
	}

	if &config.WriterMap != nil {

		switch config.WriterMap.Debug {
		case CONSOLE:
			writerMap[logrus.DebugLevel] = os.Stdout
		case FILE:
			writerMap[logrus.DebugLevel] = fileWriter
		default:
			writerMap[logrus.DebugLevel] = os.Stdout
		}

		switch config.WriterMap.Info {
		case CONSOLE:
			writerMap[logrus.InfoLevel] = os.Stdout
		case FILE:
			writerMap[logrus.InfoLevel] = fileWriter
		default:
			writerMap[logrus.InfoLevel] = os.Stdout
		}

		switch config.WriterMap.Warn {
		case CONSOLE:
			writerMap[logrus.WarnLevel] = os.Stdout
		case FILE:
			writerMap[logrus.WarnLevel] = fileWriter
		default:
			writerMap[logrus.WarnLevel] = os.Stdout
		}

		switch config.WriterMap.Error {
		case CONSOLE:
			writerMap[logrus.ErrorLevel] = os.Stdout
		case FILE:
			writerMap[logrus.ErrorLevel] = fileWriter
		default:
			writerMap[logrus.ErrorLevel] = os.Stdout
		}
	}

	// 为不同级别设置不同的输出目的
	lfHook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{})
	switch config.Level {
	case DEBUG:
		logrus.SetLevel(logrus.DebugLevel)
	case INFO:
		logrus.SetLevel(logrus.InfoLevel)
	case WARN:
		logrus.SetLevel(logrus.WarnLevel)
	case ERROR:
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.ErrorLevel)
	}

	logrus.AddHook(lfHook)
	return nil
}

// NewChannelLogger is 初始化管道类日志配置等.
func NewChannelLogger(config Logger, c chan data.JSON) error {
	writerMap := lfshook.WriterMap{}
	var fileWriter, chWriter io.Writer
	if &config.File != nil {
		_, statErr := os.Stat(config.File.Path)
		// 如果目录不存在则创建该目录
		if os.IsNotExist(statErr) {
			if err := os.MkdirAll(config.File.Path, 0755); err != nil {
				return err
			}
		}
		writer, err := rotatelogs.New(
			path.Join(config.File.Path, "%Y%m%d%H%M.log"),
			rotatelogs.WithMaxAge(time.Hour*time.Duration(config.File.MaxAgeHour)),             // 文件最大保存时间
			rotatelogs.WithRotationTime(time.Hour*time.Duration(config.File.RotationTimeHour)), // 日志切割时间间隔
		)
		if err != nil {
			return err
		}
		fileWriter = writer
	}
	writer, err := NewChanWriter(c)
	if err != nil {
		return err
	}
	chWriter = writer

	if &config.WriterMap != nil {

		switch config.WriterMap.Debug {
		case CONSOLE:
			writerMap[logrus.DebugLevel] = os.Stdout
		case FILE:
			writerMap[logrus.DebugLevel] = fileWriter
		case CHANNEL:
			writerMap[logrus.DebugLevel] = chWriter
		default:
			writerMap[logrus.DebugLevel] = os.Stdout
		}

		switch config.WriterMap.Info {
		case CONSOLE:
			writerMap[logrus.InfoLevel] = os.Stdout
		case FILE:
			writerMap[logrus.InfoLevel] = fileWriter
		case CHANNEL:
			writerMap[logrus.InfoLevel] = chWriter
		default:
			writerMap[logrus.InfoLevel] = os.Stdout
		}

		switch config.WriterMap.Warn {
		case CONSOLE:
			writerMap[logrus.WarnLevel] = os.Stdout
		case FILE:
			writerMap[logrus.WarnLevel] = fileWriter
		case CHANNEL:
			writerMap[logrus.WarnLevel] = chWriter
		default:
			writerMap[logrus.WarnLevel] = os.Stdout
		}

		switch config.WriterMap.Error {
		case CONSOLE:
			writerMap[logrus.ErrorLevel] = os.Stdout
		case FILE:
			writerMap[logrus.ErrorLevel] = fileWriter
		case CHANNEL:
			writerMap[logrus.ErrorLevel] = chWriter
		default:
			writerMap[logrus.ErrorLevel] = os.Stdout
		}
	}

	// 为不同级别设置不同的输出目的
	lfHook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{})
	switch config.Level {
	case DEBUG:
		logrus.SetLevel(logrus.DebugLevel)
	case INFO:
		logrus.SetLevel(logrus.InfoLevel)
	case WARN:
		logrus.SetLevel(logrus.WarnLevel)
	case ERROR:
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.ErrorLevel)
	}

	logrus.AddHook(lfHook)
	return nil
}

// NewEmqttLogger is 初始化消息队列类日志配置等.
func NewEmqttLogger(config Logger, emqtt mq.Emqtt) error {
	writerMap := lfshook.WriterMap{}
	var fileWriter, mqWriter io.Writer
	if &config.File != nil {
		_, statErr := os.Stat(config.File.Path)
		// 如果目录不存在则创建该目录
		if os.IsNotExist(statErr) {
			if err := os.MkdirAll(config.File.Path, 0755); err != nil {
				return err
			}
		}
		writer, err := rotatelogs.New(
			path.Join(config.File.Path, "%Y%m%d%H%M.log"),
			rotatelogs.WithMaxAge(time.Hour*time.Duration(config.File.MaxAgeHour)),             // 文件最大保存时间
			rotatelogs.WithRotationTime(time.Hour*time.Duration(config.File.RotationTimeHour)), // 日志切割时间间隔
		)
		if err != nil {
			return err
		}
		fileWriter = writer
	}

	writer, err := NewEmqttWriter(emqtt, "log")
	if err != nil {
		return err
	}
	mqWriter = writer

	if &config.WriterMap != nil {

		switch config.WriterMap.Debug {
		case CONSOLE:
			writerMap[logrus.DebugLevel] = os.Stdout
		case FILE:
			writerMap[logrus.DebugLevel] = fileWriter
		case EMQTT:
			writerMap[logrus.DebugLevel] = mqWriter
		default:
			writerMap[logrus.DebugLevel] = os.Stdout
		}

		switch config.WriterMap.Info {
		case CONSOLE:
			writerMap[logrus.InfoLevel] = os.Stdout
		case FILE:
			writerMap[logrus.InfoLevel] = fileWriter
		case EMQTT:
			writerMap[logrus.InfoLevel] = mqWriter
		default:
			writerMap[logrus.InfoLevel] = os.Stdout
		}

		switch config.WriterMap.Warn {
		case CONSOLE:
			writerMap[logrus.WarnLevel] = os.Stdout
		case FILE:
			writerMap[logrus.WarnLevel] = fileWriter
		case EMQTT:
			writerMap[logrus.WarnLevel] = mqWriter
		default:
			writerMap[logrus.WarnLevel] = os.Stdout
		}

		switch config.WriterMap.Error {
		case CONSOLE:
			writerMap[logrus.ErrorLevel] = os.Stdout
		case FILE:
			writerMap[logrus.ErrorLevel] = fileWriter
		case EMQTT:
			writerMap[logrus.ErrorLevel] = mqWriter
		default:
			writerMap[logrus.ErrorLevel] = os.Stdout
		}
	}

	// 为不同级别设置不同的输出目的
	lfHook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{})
	switch config.Level {
	case DEBUG:
		logrus.SetLevel(logrus.DebugLevel)
	case INFO:
		logrus.SetLevel(logrus.InfoLevel)
	case WARN:
		logrus.SetLevel(logrus.WarnLevel)
	case ERROR:
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.ErrorLevel)
	}

	logrus.AddHook(lfHook)
	return nil
}
