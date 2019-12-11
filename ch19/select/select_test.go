package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		retCh <- "hello"
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log("ret: ", ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
