package variable_use

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const NAME = "tomyli"

var var_name = "hello, tomyli"

var (
	name    = "name, tomyli"
	address = "address, tomyli"
)

const (
	JAVA   = "java"
	GO     = "go"
	PYTHON = "python"
)

const (
	MONDAY = iota + 1
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

const (
	ONE = 1
	TWO
	THREE
)

const (
	_ = iota
	KB
	MB
	_
	GB
	TB = iota * 2
)

func TestVarName(t *testing.T) {
	t.Log(var_name)
}

func TestVarMultiName(t *testing.T) {
	t.Log(name, address)
}

func TestLocalVarDeclare(t *testing.T) {
	var a int
	t.Log(a)
}

func TestLocalVarDeclareAndInit(t *testing.T) {
	var a = 100
	t.Log(a)
}

func TestLocalVarMultiDeclare(t *testing.T) {
	var a, b = 100, 200
	t.Log(a, b)

	var c, d = 300, "yes"
	t.Log(c, d)
}

func TestShortVarDeclare(t *testing.T) {
	a := 1
	t.Log(a)
}

func TestShortVarMultiDeclare(t *testing.T) {
	a, b := 1, 2
	t.Log(a, b)
}

func TestGlobalConstDeclare(t *testing.T) {
	t.Log(NAME)
}

func TestGlobalMultiConstDeclare(t *testing.T) {
	t.Log(JAVA, GO, PYTHON)
}

func TestLocalConstDeclare(t *testing.T) {
	const HELLO = "hello"
	t.Log(HELLO)
}

func TestGlobalConstEnumWithIota(t *testing.T) {
	t.Log(MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY)
}

func TestGlobalConstEnumDefault(t *testing.T) {
	assert.Equal(t, 1, ONE, "The two words should be the same.")
	assert.Equal(t, 1, TWO, "The two words should be the same.")
	assert.Equal(t, 1, THREE, "The two words should be the same.")
}

func TestVariableSwitch(t *testing.T) {
	a, b := 1, 2
	t.Log(a, b)

	a, b = b, a
	t.Log(a, b)
}

func TestConstWithOmit(t *testing.T) {
	t.Log(KB, MB, GB, TB)
}


func TestStructDeclareWithNew(t *testing.T) {
	timer := time.NewTimer(10)
	t.Logf("type is: %T", timer)
	t.Log(timer)
}

func TestTypeDeclare(t *testing.T) {
	type MyInt int32

	var a MyInt = 10
	var b int32 = 100

	a = MyInt(b)
	t.Log(a, b)
}

func TestTypeAlias(t *testing.T) {

	type MyInt = int32
	var a MyInt = 10
	var b int32 = 100

	a = b
	t.Log(a, b)
}