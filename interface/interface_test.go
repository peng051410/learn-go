package main_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestEmptyInterface(t *testing.T) {
	var i interface{}
	i = 1
	t.Log(i)
	i = "hello"
	t.Log(i)
}

func TestInterfaceWithNormalTypeAssertion(t *testing.T) {
	var a int64 = 10
	var i interface{} = a
	v1, ok := i.(int64)
	assert.Equal(t, int64(10), v1)
	assert.Equal(t, true, ok)
	t.Logf("v1 type is %T", v1)

	s, ok := i.(string)
	assert.Equal(t, "", s)
	assert.Equal(t, false, ok)
	t.Logf("s type is %T", s)

	v3 := i.(int64)
	assert.Equal(t, int64(10), v3)
	t.Logf("v3 type is %T", v3)

	//v4 := i.([]int64) // panic
}

type MyInterface interface {
	M1()
}

type T int

func (T) M1() {
	println("T's M1")
}

func TestInterfaceWithInterfaceTypeAssertion(t *testing.T) {

	var tt T
	var i interface{} = tt
	v1, ok := i.(MyInterface)
	assert.Equal(t, true, ok)
	assert.Equal(t, T(0), v1)
	t.Logf("v1 type is %T", v1)

	v1.M1()

	i = int64(10)
	v2, ok := i.(MyInterface)
	assert.Equal(t, false, ok)
	t.Logf("v2 type is %T", v2)
}

func TestInterfaceInfo(t *testing.T) {
	var err error
	err = errors.New("this is an error")
	t.Logf("err type is %T", err)
}

type Book interface {
	GetName() string
}

type Novel struct {
}

type Video struct {
}

func (n Novel) GetName() string {
	return "Novel"
}

func (v Video) GetName() string {
	return "Video"
}

func TestInterfaceWithDuck(t *testing.T) {

	books := []Book{new(Novel), new(Video)}
	for _, book := range books {
		t.Logf("book type is %T", book)
		t.Log(book.GetName())
	}
}

type MyError struct {
	error
}

var ErrBad = MyError{
	errors.New("this is an error"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

func TestInterfaceWithNil(t *testing.T) {
	err := returnsError()
	if err != nil {
		t.Logf("err type is %T", err)
		return
	}
	t.Log("err is nil")
}

type TT struct {
	n int
	s string
}

func TestEmptyInterfaceExpress(t *testing.T) {
	var tt = TT{
		n: 10,
		s: "hello",
	}
	var ei interface{} = tt //store with eface
	t.Logf("ei type is %T", ei)
}

type NonEmptyInterface interface {
	M1()
	M2()
}

func (TT) M1() {
	println("TT's M1")
}
func (TT) M2() {
	println("TT's M2")
}
func TestNormalInterfaceExpress(t *testing.T) {
	var tt = TT{
		n: 10,
		s: "hello",
	}
	var i NonEmptyInterface = tt //store with iface
	t.Logf("i type is %T", i)
}

func TestNilInterface(t *testing.T) {
	var i interface{}
	var err error
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i == err:", i == err)
}

func TestEmptyInterfaceType(t *testing.T) {
	var eif1 interface{}
	var eif2 interface{}
	var n, m int = 17, 18

	eif1 = n
	eif2 = m

	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 == eif2:", eif1 == eif2)

	eif2 = 17
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 == eif2:", eif1 == eif2)

	eif2 = int64(17)
	println("eif1:", eif1)
	println("eif2:", eif2)
	println("eif1 == eif2:", eif1 == eif2)
}

type TTT int

func (TTT) Error() string {
	return "a bad error"
}
func TestNonEmptyInterfaceType(t *testing.T) {
	var err1 error
	var err2 error
	err1 = (*TTT)(nil) // noempty interface type is not nil
	println("err1:", err1)
	println("err1 == nil:", err1 == nil)

	err1 = TTT(5)
	err2 = TTT(6)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 == err2:", err1 == err2)

	err2 = fmt.Errorf("%d\n", 5)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 == err2:", err1 == err2)
}

func TestNonEmptyInterfaceCompareWithEmptyInter(t *testing.T) {
	var eif interface{} = TTT(5)
	var err error = TTT(5)
	println("eif:", eif)
	println("err:", err)
	// compare with eface._type and iface.tab._type
	println("eif == err:", eif == err)

	eif = TTT(6)
	println("eif:", eif)
	println("err:", err)
	println("eif == err:", eif == err)
}

type Person struct {
	Sex
}

type Sex interface {
	Name() string
}

type Men struct {
}

func (men Men) Name() string {
	return "men"
}

type Women struct {
}

func (women Women) Name() string {
	return "women"
}

func NewPerson(sex Sex) *Person {
	return &Person{
		Sex: sex,
	}
}

func TestInterfaceCreator(t *testing.T) {
	person := NewPerson(Men{})
	t.Log(person.Name())

	person = NewPerson(Women{})
	t.Log(person.Name())
}

type capitalizedReader struct {
	r io.Reader
}

func (r capitalizedReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return 0, err
	}

	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}
	return n, err
}

func CapReader(r io.Reader) io.Reader {
	return &capitalizedReader{r: r}
}

func TestInterfaceDecorator(t *testing.T) {
	r := strings.NewReader("hello world")
	lr := io.LimitReader(r, 5)
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

	log.Println("-----------------")

	r1 := CapReader(io.LimitReader(strings.NewReader("hello world"), 5))
	if _, err := io.Copy(os.Stdout, r1); err != nil {
		log.Fatal(err)
	}
}
