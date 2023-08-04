package main

import (
	"runtime"
	"testing"
)

func TestRuntimeTraceName(t *testing.T) {
	pc, _, _, ok := runtime.Caller(0)
	if ok {
		name := runtime.FuncForPC(pc).Name()
		t.Log("name is: ", name)
	}
}
