package generics_test

import (
	"fmt"
	"testing"
)

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func TestSumIntsOrFloats(t *testing.T) {
	ints := map[int]int64{1: 1, 2: 2, 3: 3}
	t.Log(SumIntsOrFloats(ints))
	t.Log(SumIntsOrFloats(map[int]float64{1: 1.1, 2: 2.2, 3: 3.3}))
}

// test interface define type constraint
type C1 interface {
	~int | ~int32
	M1()
}

type T struct {
}

func (t *T) M1() {
}

type T1 int

func (t T1) M1() {}

func foo[P C1](t P) {

}

func TestInterfaceTypeConstraint(t *testing.T) {
	var t1 T1
	foo(t1)
	// foo(T{}) // T does not implement C1 (missing M1 method)
	// foo("abc") // error
}

func Sort[Elem interface{ Less(y Elem) bool }](list []Elem) {

}

type book struct{}

func (x book) Less(y book) bool {
	return true
}

func TestTypeInstantiation(t *testing.T) {
	var bookshelf []book
	// Sort[book](bookshelf)
	Sort(bookshelf) //compiler can infer the type
}

type Vector[T any] []T

func (v Vector[T]) Dump() {
	fmt.Printf("%#v\n", v)
}

func TestGenericsType(t *testing.T) {
	var iv = Vector[int]{1, 2, 3, 4}
	var sv Vector[string]
	sv = []string{"a", "b", "c"}
	iv.Dump()
	sv.Dump()
}

type Orm interface {
	// Insert[T any](data ...T)(sql.Result, error) //interface method must have no type parameters
}

func maxAny(sl []any) any {
	if len(sl) == 0 {
		return nil
	}
	max := sl[0]
	for _, v := range sl[1:] {
		switch v.(type) {
		case int:
			if v.(int) > max.(int) {
				max = v
			}
		case string:
			if v.(string) > max.(string) {
				max = v
			}
		case float64:
			if v.(float64) > max.(float64) {
				max = v
			}
		}
	}
	return max
}

func TestMaxAny(t *testing.T) {
	t.Log(maxAny([]any{1, 2, 3, 4, 5}))
	t.Log(maxAny([]any{"a", "b", "c", "d"}))
	t.Log(maxAny([]any{1.1, 2.2, 3.3, 4.4}))
}

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 |
		~string
}

func maxGeneric[T ordered](sl []T) T {
	if len(sl) == 0 {
		panic("empty input")
	}
	max := sl[0]
	for _, v := range sl[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

type mystring string

func TestMaxGeneric(t *testing.T) {
	// int is type real param
	t.Log(maxGeneric([]int{1, 2, 3, 4, 5}))
	t.Log(maxGeneric([]string{"a", "b", "c", "d"}))
	t.Log(maxGeneric([]float64{1.1, 2.2, 3.3, 4.4}))
	t.Log(maxGeneric([]mystring{"a", "b", "c", "d"}))

	var intArr = []int{1, 2, 3, 4, 5}
	// compiler can infer the type
	t.Log(maxGeneric(intArr))

	maxGenericInt := maxGeneric[int]
	t.Logf("%T", maxGenericInt)
	t.Log(maxGenericInt(intArr))
}
