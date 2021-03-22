package log

import (
	"io"
	"log"
	"os"
)

var (
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(format string, v ...interface{}) {
	info.Printf(format, v)
}

func Warning(format string, v ...interface{}) {
	warning.Printf(format, v)
}

func Error(format string, v ...interface{}) {
	error.Printf(format, v)
}
