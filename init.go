package gotask

import (
	"sync"
	"time"
)

//任务数据结构
type task struct {
	timestamp int64             //unix时间戳
	taskFunc  func(interface{}) //执行的函数
	param     interface{}       //函数的参数
}

//任务map,[任务名]=>task结构
var taskMap = sync.Map{}

func init() {
	go taskLoop()
}

//任务执行循环
func taskLoop() {
	//搜索到期执行的任务
	for {
		now := time.Now().Unix()
		taskMap.Range(func(name, t interface{}) bool {
			t2 := t.(task)
			n := name.(string)

			//未到时间，不执行
			if t2.timestamp < now {
				return true
			}

			//已到执行时间，执行任务
			go runTask(n, &t2)
			return true
		})

		time.Sleep(time.Second)
	}
}

//执行任务,记录日志
func runTask(tname string, t *task) {
	t.taskFunc(t.param)

}
