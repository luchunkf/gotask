package gotask

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger
var file *os.File

//根据path设置日志
func setLogger(path string) {

	var err error
	file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// buff := bufio.NewWriter(file)

	logger = log.New(file, "", log.LstdFlags)
}

//写入一条执行成功日志
func logRunSuccess(name string) {
	if logger == nil {
		return
	}

	logger.Println("任务: " + name + "执行成功!")
}

//写入一条执行失败日志
func logRunFail(name string, err interface{}) {

	if logger == nil {
		return
	}

	logger.Printf(`任务: %s执行失败!
	失败原因:
	%s
	`, name, err)
}
