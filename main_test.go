package gotask

import (
	"fmt"
	"testing"
	"time"
)

func task2(v interface{}) {
	fmt.Println("run task2")
	fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"))
	fmt.Println(ListTask())
}

func Test1(t *testing.T) {

	SetLogPath("./task-log.log")

	runTime := time.Now().Add(time.Second * 2)
	AddTask("task1", runTime, func(v interface{}) {
		fmt.Println("run task1")
		fmt.Println(v)
		fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"))
		fmt.Println(ListTask())
	}, "abc")

	AddTask("task2", time.Now().Add(time.Second*8), task2, nil)

	AddTask("task3", time.Now().Add(time.Second*4), func(v interface{}) {
		fmt.Println(ListTask())
		s := make([]int, 1)
		fmt.Println(s[100])
	}, nil)

	time.Sleep(time.Second * 10)
}

func Example1() {

	SetLogPath("./task-log.log")

	runTime := time.Now().Add(time.Second * 2)
	AddTask("task1", runTime, func(v interface{}) {
		fmt.Println("run task1")
		fmt.Println(v)
		fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"))
		fmt.Println(ListTask())
	}, "abc")

	AddTask("task2", time.Now().Add(time.Second*8), task2, nil)

	AddTask("task3", time.Now().Add(time.Second*4), func(v interface{}) {
		fmt.Println(ListTask())
		s := make([]int, 1)
		fmt.Println(s[100])
	}, nil)

	time.Sleep(time.Second * 10)
}
