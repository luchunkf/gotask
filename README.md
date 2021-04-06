# gotask

## gotask介绍
the go package, use to run task in specified time
用于go语言，定时执行任务，特点:
- 线程安全
- 支持高并发执行任务
- 定时执行任务
- 支持运行时获取任务列表
- 支持运行时撤销任务

---

## gotask使用

### 获取：
`go get -u github.com/luchunkf/gotask`

### 导入：  
`import "github.com/luchunkf/gotask"`

### 添加任务

1.添加一段时间后执行的任务

```go
	t := time.Now().Add(time.Second * 5)
	gotask.AddTask("task1", t, func(v interface{}) {
        //do some thing
	}, nil)
``` 

2.添加指定时间后执行的任务

```go
	t := time.Now().Add(time.Second * 5)
	gotask.AddTask("task1", t, func(v interface{}) {
        //do some thing
	}, nil)
``` 