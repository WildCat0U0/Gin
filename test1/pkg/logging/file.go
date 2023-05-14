package logging

import (
	"fmt"
	"log"
	"os"
)

// 中文
var (
	LogSavePath = "test1/runtime/logs/" //日志保存路径
	LogSaveName = "log"                 //日志文件名
	LogFileExt  = "log"                 //日志文件后缀
	TimeFormat  = "20060102"            //时间格式
)

// 说明一下这个函数的作用
// 1.定义了一个getLogFileFullPath函数，用来返回日志保存路径+日志文件名
// 2.定义了一个openLogFile函数，用来打开日志文件
// 3.定义了一个mkDir函数，用来创建日志文件夹
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath) //返回日志保存路径
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()                                            //返回日志保存路径
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, TimeFormat, LogFileExt) //返回日志文件名
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)                        //返回日志保存路径+日志文件名
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath) //返回文件信息
	switch {
	case os.IsNotExist(err): //判断是否存在
		mkDir() //创建文件夹
	case os.IsPermission(err): //判断是否有权限
		log.Fatalf("Permission :%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //打开文件
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err) //打开失败
	}
	return handle
}

func mkDir() {
	dir, _ := os.Getwd()                                      //获取当前路径
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm) //创建文件夹
	if err != nil {
		panic(err) //创建失败
	}
}

//os.Getwd：返回与当前目录对应的根路径名
//os.MkdirAll：创建对应的目录以及所需的子目录，若成功则返回nil，否则返回error
//os.ModePerm：const定义ModePerm FileMode = 0777
