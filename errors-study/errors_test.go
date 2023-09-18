package errors_study_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError(t *testing.T) {
	err := errors.New("this is an error")
	t.Log(err)
}

func TestErrorFromFmt(t *testing.T) {

	err := fmt.Errorf("this is an error %s", "tomyli")
	t.Log(err)
}

func TestGetErrorMessage(t *testing.T) {

	err := fmt.Errorf("this is an error %s", "tomyli")
	t.Log(err.Error())
}

func TestErrorsIsFunc(t *testing.T) {

	err := errors.New("this is an error")
	err1 := fmt.Errorf("this is an error %w", err)
	err2 := fmt.Errorf("this is an error %w", err1)

	assert.Equal(t, true, errors.Is(err1, err), "The two words should be the same.")
	assert.Equal(t, true, errors.Is(err2, err), "The two words should be the same.")
}

type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

func TestErrorAsFunc(t *testing.T) {
	err := &MyError{"this is an error"}
	err1 := fmt.Errorf("wrap error: %w", err)
	err2 := fmt.Errorf("wrap error: %w", err1)
	var e *MyError
	assert.Equal(t, true, errors.As(err2, &e), "find first error fo MyError")
	assert.Equal(t, true, e == err, "same errors value")

	// normal use case
	//if errors.As(err2, &e) {
	//	t.Log(e == err)
	//}
}

func foo() {
	println("foo start")
	panic("no no no")
	bar()
	println("foo end")
}

func TestPanicBroadcast(t *testing.T) {
	println("test start")
	foo()
	println("test end")
}

func TestPanicBroadcastNormal(t *testing.T) {
	println("test start")
	fooWithRecover()
	println("test end")
}

func fooWithRecover() {
	defer func() {
		if r := recover(); r != nil {
			println("got panic")
		}
	}()
	println("foo start")
	panic("no no no")
	bar()
	println("foo end")
}

func bar() {
	println("bar start")
}

func TestPanicRecover(t *testing.T) {
	assertPanic(t, foo)
}

// https://stackoverflow.com/questions/31595791/how-to-test-panics
func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestDeferFunc(t *testing.T) {
	defer func() {
		t.Log("clear value")
	}()

	t.Log("do job")
}

func TestDeferFuncClear(t *testing.T) {
	defer clear()
	t.Log("do job")
}

func clear() {
	println("do clear")
}

func TestDeferOrder(t *testing.T) {

	defer func() {
		t.Log("one")
	}()

	defer func() {
		t.Log("two")
	}()

	t.Log("hello")
}

func TestDeferValidTime(t *testing.T) {

	panic("I'm scared")
	t.Log("do do do")
	defer clear()
	t.Log("run success")
}

func TestDeferFuncVariableBind(t *testing.T) {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}

	t.Log("====================")
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			t.Log(n)
		}(i)
	}

	t.Log("====================")

	for i := 0; i <= 3; i++ {
		defer func() {
			t.Log("unexpect value", i)
		}()
	}

}

func c() (i int) {
	defer func() {
		i++
	}()
	return 1
}
func TestDeferFuncAssignReturnValue(t *testing.T) {
	t.Log(c())
}
