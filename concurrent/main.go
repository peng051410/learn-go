package main

import (
	"fmt"
	"time"
)

func main() {

	count := 10
	for i := 0; i < count; i++ {
		go func (i int)  {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 1)
}