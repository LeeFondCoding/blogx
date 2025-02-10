package core

import (
	"blogx/global"
	
	"bytes"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

// 日志的颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}


func (*LogFormatter) Format(entry *logrus.Entry)([]byte, error) {
	// 每个日志级别都有不同的颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var buf *bytes.Buffer
	if entry.Buffer != nil {
		buf = entry.Buffer
	} else {
		buf = &bytes.Buffer{}
	}

	formatTime := entry.Time.Format("2006-01-21 15:09:14")
	
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		fmt.Fprintf(buf, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", formatTime, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(buf, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", formatTime, levelColor, entry.Level, entry.Message)
	}
	return buf.Bytes(), nil
}

type FileDateHook struct {
	file *os.File
	logPath string
	fileDate string // 判断日期切换目录
	appName string
}

func (hook *FileDateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *FileDateHook) Fire(entry *logrus.Entry) error {
	formatTime := entry.Time.Format("2006-01-02")
	//fmt.Println(formatTime)
	line, _ := entry.String()
	if hook.fileDate == formatTime {
		hook.file.Write([]byte(line))
		return nil
	}
	hook.file.Close()
	os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, formatTime), os.ModePerm)
	fileName := fmt.Sprintf("%s/%s/%s.log", hook.logPath, formatTime, hook.appName)

	hook.file, _ = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = formatTime
	hook.file.Write([]byte(line))
	return nil
}

func InitFile(logPath, appName string) {
	fileDate:= time.Now().Format("2006-01-02")
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return
	}

	filename := fmt.Sprintf("%s/%s/%s.log", logPath, fileDate, appName)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		logrus.Error(err)
		return
	}
	fileHook := FileDateHook{file, logPath, fileDate, appName}
	logrus.AddHook(&fileHook) 
}

func InitLogrus() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetLevel(logrus.TraceLevel)
	l := global.Conf.Log
	InitFile(l.Dir, l.App)
}
