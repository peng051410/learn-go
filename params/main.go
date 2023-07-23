package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("type is %T", os.Args)
	fmt.Println(os.Args)
}
