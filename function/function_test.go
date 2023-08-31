package function_test

import (
	"math/rand"
	"testing"
	"time"
)

func TestFunctionAsVariable(t *testing.T) {

	a := func(name string) {
		t.Log("hello", name)
	}
	a("tomyli")
}

func bigOne() func() int {
	return func() int {
		return 10
	}
}

func TestFunctionReturnFunction(t *testing.T) {
	i := bigOne()()
	t.Log(i)
}

func TestFunctionAsParameter(t *testing.T) {
	time.AfterFunc(1*time.Second, func() {
		t.Log("after func")
	})
}

func TestCustomDefineFunctionType(t *testing.T) {
	param := doAsParam(doTransfer)
	t.Log(param)
}

func FuncClosure(num int32) func() int32 {
	return func() int32 {
		return num + 5
	}
}

func TestFunctionClosure(t *testing.T) {
	f := FuncClosure(10)
	t.Log(f())
}

func doTransfer(name string) string {
	return "hello " + name
}

func doAsParam(f func(string) string) string {
	return f("tomyli")
}

func returnMultiValues() (int, int) {
	return rand.Int(), rand.Int()
}

func TestFunctionMultiReturnValues(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}
