package type_study

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

type MyInt int64
type MyInitAlias = int64

func TestIntTypeDeclare(t *testing.T) {

	var a int = -1
	var b int8 = 1
	var c int16 = 1
	var d int32 = 1
	var e int64 = 1

	f := 1

	t.Log("a: ", unsafe.Sizeof(a))
	t.Log("b: ", unsafe.Sizeof(b))
	t.Log("c: ", unsafe.Sizeof(c))
	t.Log("d: ", unsafe.Sizeof(d))
	t.Log("e: ", unsafe.Sizeof(e))
	t.Log("f: ", unsafe.Sizeof(f))
	t.Logf("f type is %T: ", f)
}

func TestUIntTypeDeclare(t *testing.T) {

	var a uint = 1
	var b uint8 = 1
	var c uint16 = 1
	var d uint32 = 1
	var e uint64 = 1
	//var f uint = -1 // cannot convert -1 (untyped int constant) to uint

	t.Log("a: ", unsafe.Sizeof(a))
	t.Log("b: ", unsafe.Sizeof(b))
	t.Log("c: ", unsafe.Sizeof(c))
	t.Log("d: ", unsafe.Sizeof(d))
	t.Log("e: ", unsafe.Sizeof(e))

	var g uintptr = 0xc0000b4000
	var h uintptr = uintptr(unsafe.Pointer(&g))
	t.Log(g)
	t.Logf("g type is %T", g)
	t.Log(h)
	t.Logf("h type is %T", h)
}

func TestBoolDeclare(t *testing.T) {
	var a bool = true
	var b bool = false
	t.Log("a: ", a)
	t.Log("b: ", b)
}

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	//b = a // cannot use a (type int32) as type int64 in assignment
	b = int64(a)

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1 //invalid operation: aPtr + 1 (mismatched types *int and int)
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
	t.Log(*aPtr)
}

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestCustomFn(t *testing.T) {
	a := timeSpent(slowFun)
	t.Log(a(10))
}
