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
