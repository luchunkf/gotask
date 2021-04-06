package gotask

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//添加任务
func AddTask(name string, t time.Time, f func(interface{}), param interface{}) {

	//构造唯一任务名
	timeStr := t.Local().Format("2006-01-02_15:04:05")
	randStr := strconv.Itoa(rand.Intn(1000000))
	uName := fmt.Sprintf("%s-%s-%s", name, timeStr, randStr)

	taskMap.Store(uName, task{
		timestamp: t.Unix(),
		taskFunc:  f,
		param:     param,
	})
}

//删除任务
func DelTask(uName string) error {
	_, has := taskMap.LoadAndDelete(uName)
	if !has {
		return errors.New("no found task")
	}
	return nil
}

//获取任务列表
func ListTask() []string {

	list := make([]string, 0)

	taskMap.Range(func(name, t interface{}) bool {
		n := name.(string)
		list = append(list, n)
		return true
	})

	return list

}

//设置日志Path
func SetLogPath(path string) {
	if file != nil {
		file.Close()
	}
	setLogger(path)
}
