package logs

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 实现一个自定义的logrus.Formatter
type LogFormatter struct {
}

// Format 方法名称是固定，不能修改
func (m *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 注意： logrus 的默认 Level 设置是 Info 级别

	levelColor := color.New(color.FgWhite).Add(color.Bold)
	switch entry.Level {
	case logrus.PanicLevel:
		levelColor = color.New(color.FgRed).Add(color.Bold)
	case logrus.FatalLevel:
		levelColor = color.New(color.FgRed).Add(color.Bold)
	case logrus.ErrorLevel:
		levelColor = color.New(color.FgRed)
	case logrus.WarnLevel:
		levelColor = color.New(color.FgYellow)
	case logrus.InfoLevel:
		levelColor = color.New(color.FgCyan)
	case logrus.DebugLevel:
		levelColor = color.New(color.FgBlue)
	}

	// 格式化时间、级别和消息
	timeStr := levelColor.SprintFunc()(entry.Time.Format("2006-01-02 15:04:05.000"))
	levelStr := levelColor.SprintFunc()(entry.Level.String())

	newLog := fmt.Sprintf("[%s] [%s]\n%s\n", timeStr, levelStr, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func InitConfig() {
	// 设置日志格式
	// logrus.SetFormatter(&logrus.TextFormatter{
	// 	ForceColors:     true,
	// 	DisableColors:   true, // 可根据需要禁用颜色
	// 	ForceQuote:      true, // 关闭强制引号
	// 	FullTimestamp:   true,
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })

	// logrus.SetFormatter(&logrus.JSONFormatter{
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })

	fmt.Println("logrus init")

	// 注意： logrus 的默认 Level 设置是 Info 级别
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&LogFormatter{})

	// 设置日志文件
	log := &lumberjack.Logger{
		Filename:   "./logs/gin.log", // 日志文件路径
		MaxSize:    10,               // 文件大小限制，单位MB
		MaxBackups: 3,                // 最多保留几个备份
		MaxAge:     15,               // 文件最多保存多少天
		Compress:   true,             // 是否压缩/归档旧文件
	}

	// 创建控制台输出器
	consoleWriter := logrus.New()
	consoleWriter.Out = os.Stdout
	consoleWriter.Level = logrus.DebugLevel
	consoleWriter.Formatter = &LogFormatter{}

	multi := io.MultiWriter(consoleWriter.Out, log)

	// 设置全局日志输出为文件
	// logrus.SetOutput(log)
	logrus.SetOutput(multi)
}
