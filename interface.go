package gotask

import (
	"errors"
	"time"
)

//添加任务
func AddTask(name string, t time.Time, f func(interface{}), param interface{}) {
	taskMap.Store(name, task{
		timestamp: t.Unix(),
		taskFunc:  f,
		param:     param,
	})
}

//删除任务
func DelTask(name string) error {
	_, has := taskMap.LoadAndDelete(name)
	if !has {
		return errors.New("no found task")
	}
	return nil
}
