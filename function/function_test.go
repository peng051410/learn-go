package function_test

import (
	"testing"
	"time"
)

func TestFunctionAsVariable(t *testing.T) {

	a := func(name string) {
		t.Log("hello", name)
	}
	a("tomyli")
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

func TestFunctionClosure(t *testing.T) {
}

func doTransfer(name string) string {
	return "hello " + name
}

func bigOne() func() int {
	return func() int {
		return 10
	}
}

func doAsParam(f func(string) string) string {
	return f("tomyli")
}
