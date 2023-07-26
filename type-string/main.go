package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	showStringOriginalData()
}

func showStringOriginalData() {
	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("0x%x\n", hdr.Data)
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))
	dumpBytesArray((*p)[:])
}

func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c,", b)
	}
	fmt.Printf("]\n")
}
