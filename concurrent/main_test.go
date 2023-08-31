package main_test

import (
	"fmt"
	"testing"
	"time"
)

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func service() string {
	time.Sleep(time.Millisecond * 100)
	return "service done."
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	fmt.Println("do other thing.")
	time.Sleep(time.Second * 1)
	fmt.Println(<-retCh)
}
