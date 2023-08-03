package method_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type Person struct {
	Name string
}

func (p Person) Print() string {
	return "the person name is " + p.Name
}

func (p Person) Get() string {
	return p.Name
}

func (p *Person) Set(name string) {
	p.Name = name
}

func TestMethodDefine(t *testing.T) {

	person := Person{}
	person.Name = "tomyli"
	t.Log(person.Print())
}

func TestMethodInvoke(t *testing.T) {
	person := Person{
		Name: "tomyli",
	}
	t.Log(person.Print())

	p2 := &Person{
		Name: "tomyli2",
	}
	t.Log(p2.Print())
}

func TestMethodGetAndSet(t *testing.T) {
	person := Person{
		Name: "tomyli",
	}
	t.Log(person.Get())

	person.Set("tomyli2")
	t.Log(person.Get())
}

func TestMethodGetAndSetSpread(t *testing.T) {
	person := Person{
		Name: "tomyli",
	}
	t.Log(person.Get())

	(&person).Set("tomyli2")
	t.Log(person.Get())
}

func TestMethodExpression(t *testing.T) {
	person := Person{
		Name: "tomyli",
	}
	t.Log(Person.Get(person))

	(*Person).Set(&person, "tomyli2")

	t.Log(person.Get())
}

func TestMethodExpressionAsValue(t *testing.T) {
	person := Person{
		Name: "tomyli",
	}
	f1 := (*Person).Set
	f2 := Person.Get
	fmt.Printf("%T, %T\n", f1, f2)
	f1(&person, "tomyli2")
	t.Log(f2(person))
}

type T struct {
	a int
}

func (t T) F1() {
	t.a = 10
}

func (t *T) F2() {
	t.a = 11
}

func TestMethodReceiver(t *testing.T) {
	var at T
	t.Log(at.a)

	at.F1()
	assert.Equal(t, 0, at.a, "the value should be 0")

	p := &at
	p.F2()
	assert.Equal(t, 11, p.a, "the value should be 11")
}

func TestMethodReceiverTypeInvoke(t *testing.T) {

	var t1 T
	t.Log(t1.a)
	assert.Equal(t, 0, t1.a, "the value should be 0")
	t1.F1()
	t.Log(t1.a)
	assert.Equal(t, 0, t1.a, "the value should be 0")
	t1.F2() //go compile auto convert to (&t1).F2()
	t.Log(t1.a)
	assert.Equal(t, 11, t1.a, "the value should be 11")

	var t2 = &T{}
	t.Log(t2.a)
	assert.Equal(t, 0, t2.a, "the value should be 0")
	t2.F1() //go compile auto convert to (*t2).F1()
	t.Log(t2.a)
	assert.Equal(t, 0, t2.a, "the value should be 0")
	t2.F2()
	t.Log(t2.a)
	assert.Equal(t, 11, t2.a, "the value should be 11")
}

type Interface interface {
	M1()
	M2()
}

type T1 struct{}

func (t T1) M1() {
}
func (t T1) M3() {
}

func (t *T1) M2() {
}
func (t *T1) M4() {
}

func TestMethodSet(t *testing.T) {

	var t1 T1
	var t2 *T1
	t.Log(t1)
	//var i1 Interface = t1 // compile error t1 is not implement Interface M2()
	var i2 Interface = t2

	i2.M1()
	i2.M2()
}

func TestMethodSetInfo(t *testing.T) {

	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t1 T1
	var t2 *T1
	dumpMethodSet(t1)
	dumpMethodSet(t2)
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

type T2 struct {
}

func (t T2) M1() {
}
func (t T2) M2() {
}

type S T2

func TestMethodSetTypeDefine(t *testing.T) {

	var s S
	dumpMethodSet(s)
	dumpMethodSet(&s)

	b := T2(s)
	dumpMethodSet(b)
	dumpMethodSet(&b)
}
