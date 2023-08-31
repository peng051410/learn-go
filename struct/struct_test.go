package struct_test

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestStructDeclare(t *testing.T) {
	type User struct {
		Name    string
		Age     int
		Address string
	}

	user := User{}
	t.Log("age is: ", user.Age)
}

func TestStructDeclareEmpty(t *testing.T) {
	type User struct{}
	var user User
	t.Log("size of struct is: ", unsafe.Sizeof(user))
}

func TestStructWithOtherStruct(t *testing.T) {
	type Address struct {
		Name string
		Code string
	}

	type User struct {
		Name    string
		Age     int
		Address Address
	}

	user := User{}
	t.Log("code is: ", user.Address.Name)
	//t.Log("code is: ", user.Code)
}

func TestStructWithImportOtherStruct(t *testing.T) {
	type Address struct {
		Name string
		Code int
	}

	type User struct {
		Name string
		Age  int
		//匿名引入
		Address
	}

	user := User{}
	t.Log("code is: ", user.Address.Code)
	t.Log("code is: ", user.Code)
}

func TestStructWithPointer(t *testing.T) {
	type User struct {
		up *User
		us []User
		um map[string]User
	}

	user := User{}
	t.Log("user is: ", user)
}

func TestStructInit(t *testing.T) {

	type User struct {
		Name string
		Age  int
	}

	user := User{
		Name: "John",
		Age:  1,
	}
	t.Log(user)

	var badUser User
	badUser.Name = "Tom"
	badUser.Age = 0
	t.Log(badUser)
}

func TestStructZeroInit(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	// Complie success, but no meaning
	var user User
	t.Log(user)
}

func TestStructDeclareBadCase(t *testing.T) {
	type User struct {
		Name string
		// add a new name, construct init complie error
		// Nick string
		Age int
	}

	user := User{"TOM", 11}
	t.Log(user)
}

type MyInt int

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}

type t struct {
	a int
	b int
}

type S struct {
	*MyInt
	t
	io.Reader
	s string
	n int
}

func TestStructEmbedded(tt *testing.T) {
	m := MyInt(17)
	r := strings.NewReader("hello, tomy")
	s := S{
		MyInt:  &m,
		t:      t{a: 1, b: 2},
		Reader: r,
		s:      "hello",
	}

	var sl = make([]byte, len("hello, tomy"))
	s.Reader.Read(sl)
	s.Read(sl)
	fmt.Println(string(sl))
	s.MyInt.Add(18)
	s.Add(18)
	fmt.Println(*(s.MyInt))
}

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Printf("nil\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s no method\n", dynTyp)
		return
	}
	fmt.Printf("%s has %d methods:\n", dynTyp, n)
	for i := 0; i < n; i++ {
		method := dynTyp.Method(i)
		fmt.Println("-", method.Name)
	}
	fmt.Printf("\n")
}

type I interface {
	M1()
	M2()
}
type T struct {
	I
}

func (T) M3() {}

func TestStructEmbeddedInterface(t *testing.T) {
	var t1 T
	dumpMethodSet(t1)
	dumpMethodSet(&t1)
}

type T1 struct {
}

func (T1) T1M1()   {}
func (*T1) PT1M2() {}

type T2 struct {
}

func (T2) T2M1()   {}
func (*T2) PT2M2() {}

type T3 struct {
	T1
	*T2
}

func TestStructEmbeddedInterface2(t *testing.T) {
	tt := T3{
		T1: T1{},
		T2: &T2{},
	}

	dumpMethodSet(tt)
	dumpMethodSet(&tt)
}

type A struct {
}

func (A) M1()  {}
func (*A) M2() {}

type A1 A
type A2 = A

func TestStructEmbeddedDefineType(t *testing.T) {

	var a A
	var pt *A
	var a1 A1
	var pt1 *A1

	dumpMethodSet(a)
	dumpMethodSet(pt)

	dumpMethodSet(a1)
	dumpMethodSet(pt1)
}

func TestStructEmbeddedAliasType(t *testing.T) {

	var a A
	var pt *A
	var a2 A2
	var pt2 *A2

	dumpMethodSet(a)
	dumpMethodSet(pt)

	dumpMethodSet(a2)
	dumpMethodSet(pt2)
}
