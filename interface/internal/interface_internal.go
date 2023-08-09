package main

import "fmt"

type TT struct {
	n int
	s string
}

func (TT) M1() {}
func (TT) M2() {}

type NonEmptyInterface interface {
	M1()
	M2()
}

func main() {
	var t = TT{
		n: 10,
		s: "hello",
	}
	var ei interface{}
	ei = t

	var i NonEmptyInterface
	i = t
	fmt.Println(ei)
	fmt.Println(i)
}
