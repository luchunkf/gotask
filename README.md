# gotask

## gotask介绍

the go package, use to run task in specified time
用于go语言，定时执行任务，特点:

- 线程安全
- 支持高并发执行任务
- 定时执行任务
- 支持运行时获取任务列表
- 支持运行时撤销任务
- 支持执行日志

---

## gotask使用

### 获取：
`go get -u github.com/luchunkf/gotask`

### 导入：  
`import "github.com/luchunkf/gotask"`

### 添加任务

函数说明:  

```go
gotask.AddTask(name string, t time.Time, f func(interface{}), param interface{})
```

参数：

- name: 任务名  
- t: 任务执行时间
- f: 任务执行的函数，必须满足func(interface{})签名
- param: 任务执行时所传入的参数(如果没有参数可以传入nil)
  
> 为了避免任务名冲突，每个任务都会生成一个**实际的任务**名。规则`name_执行时间_随机数`, 如: `task2-2021-04-06_22:26:56-727887`  


1. 添加一段时间后执行的任务

```go
t := time.Now().Add(time.Second * 5)
gotask.AddTask("task1", t, func(v interface{}) {
    //do some thing
}, "abc")
```

2. 添加指定时间后执行的任务

```go
t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-04-06 22:44:00", time.Local)
gotask.AddTask("task1", t, func(v interface{}) {
    //do some thing
}, nil)
```

> 如果添加的任务时间小于当前时间，则任务在下一秒内立即执行

---  

### 获取任务列表

函数说明:  

```go
gotask.ListTask() []string
```

获取当前已加入执行队列中的**实际任务名**切片，例如:`[task2-2021-04-06_22:26:56-727887 task3-2021-04-06_22:26:52-131847]`
> 已经执行的任务会从任务队列中移除

### 删除任务

函数说明:  

```go
gotask.DelTask(uName string) error

```

### 任务执行日志

函数说明:  

```go
gotask.SetLogPath(path string)
```

参数：

- path: 日志文件路径
  
> 如果不设置日志文件路径，则不输出日志。日志路径可以运行时随时修改
例：
```go
gotask.SetLogPath("./go_task.log") //在程序执行目录go_task.log文件记录日志
```

日志内容，例:
```
2021/04/07 13:18:56 任务: task1-2021-04-07_13:18:56-498081执行成功! //执行成功
2021/04/07 13:18:58 任务: task3-2021-04-07_13:18:58-131847执行失败! //执行失败，打印出失败原因
	失败原因:
	runtime error: index out of range [100] with length 1
	
2021/04/07 13:19:02 任务: task2-2021-04-07_13:19:02-727887执行成功!
```

### 注意事项

- 暂时不支持任务持久化，一旦主进程退出，队列中任务不会保存
- 在任务中如果出现没有被recover的panic, 并不会导致程序退出，会在日志中记录。
- 任务执行时间和实际时间误差不超过1秒。
