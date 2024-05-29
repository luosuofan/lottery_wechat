package config

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

type logFormatter struct {
	logrus.TextFormatter
}

// 自定义打印格式
func (f *logFormatter) Format(entry *logrus.Entry) ([]byte, error) { //entry包含了堆栈所有信息
	prettyCaller := func(frame *runtime.Frame) string {
		_, fileName := filepath.Split(frame.File)
		return fmt.Sprintf("%s,%d", fileName, frame.Line) //返回文件名和对应行数
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	b.WriteString(fmt.Sprintf("[%s] %s", entry.Time.Format(f.TimestampFormat), strings.ToUpper(entry.Level.String()))) //打印日志时间
	if entry.HasCaller() {                                                                                             //如果有调用者,打印堆栈信息
		b.WriteString(fmt.Sprintf("[%s]", prettyCaller(entry.Caller))) //打印文件名和对应行数
	}
	b.WriteString(fmt.Sprintf("[%s]", entry.Message)) //如果有调用者（entry.Buffer != nil）
	return b.Bytes(), nil
}
