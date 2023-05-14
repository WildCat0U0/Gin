package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

//说明一下这个文件的作用
//1.定义了一个Level类型，用来表示日志级别
//2.定义了一个F变量，用来存放日志文件的指针
//3.定义了一个logger变量，用来作为日志的输出对象
//4.定义了一个logPrefix变量，用来存放日志的前缀
//5.定义了一个levelFlags切片，用来存放日志级别的名字
//6.定义了一个init函数，用来初始化日志文件
//7.定义了一个getLogFileFullPath函数，用来返回日志保存路径+日志文件名
//8.定义了一个openLogFile函数，用来打开日志文件
//9.定义了一个setPrefix函数，用来设置日志前缀
//10.定义了一个Debug函数，用来打印Debug级别的日志
//11.定义了一个Info函数，用来打印Info级别的日志
//12.定义了一个Warn函数，用来打印Warn级别的日志
//13.定义了一个Error函数，用来打印Error级别的日志
//14.定义了一个Fatal函数，用来打印Fatal级别的日志
//15.定义了一个getLogFilePath函数，用来返回日志保存路径
//16.定义了一个getLogFileFullPath函数，用来返回日志保存路径+日志文件名
//17.定义了一个openLogFile函数，用来打开日志文件
//18.定义了一个mkDir函数，用来创建日志文件夹
//19.定义了一个getCallerInfo函数，用来获取调用日志函数的文件名、行号、函数名

type Level int

var (
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// getLogFileFullPath：返回日志保存路径+日志文件名
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// setPrefix：设置日志前缀
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING) //设置日志前缀
	logger.Println(v)  //打印日志
}
func Error(v ...interface{}) {
	setPrefix(ERROR)  //设置日志前缀
	logger.Println(v) //打印日志
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
