package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var goroutineSpace = []byte("goroutine ")
var mu = sync.Mutex{}
var m = make(map[uint64]int)

func printTrace(id uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += " "
	}
	fmt.Printf("g[%05d]: %s%s%s\n", id, indents, arrow, name)
}

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("can't get the caller info")
	}

	name := runtime.FuncForPC(pc).Name()
	gid := http2curGoroutineID()

	mu.Lock()
	indents := m[gid]
	m[gid] = indents + 1
	mu.Unlock()
	printTrace(gid, name, "-->", indents+1)

	return func() {
		mu.Lock()
		indents := m[gid]
		m[gid] = indents - 1
		mu.Unlock()
		printTrace(gid, name, "<--", indents)
	}
	//fmt.Printf("goroutine id: [%05d]: enter [%s]\n", gid, name)
	//return func() {
	//	fmt.Printf("goroutine id: [%05d]: exit [%s]\n", gid, name)
	//}
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}

func A1() {
	defer Trace()()
	A2()
}

func A2() {
	defer Trace()()
	A3()
}

func A3() {
	defer Trace()()
}

func B1() {
	defer Trace()()
	B2()
}

func B2() {
	defer Trace()()
	B3()
}

func B3() {
	defer Trace()()
}

func main() {
	//defer Trace()()
	//foo()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		B1()
		wg.Done()
	}()

	A1()
	wg.Wait()
}

func http2curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}
