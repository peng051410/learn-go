package type_study

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
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
