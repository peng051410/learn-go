package condition_test

import (
	"testing"
)

func TestConditionIf(t *testing.T) {
	a, b := 1, 2
	if b != a {
		t.Log("a != b")
	}
}

func TestConditionWithVariable(t *testing.T) {
	b := 2
	if c := 10; c > b {
		t.Log("c > b")
	} else if d := 20; d > b {
		t.Log("d > b")
	}
}

func TestFor(t *testing.T) {

	for i := 0; i < 5; i++ {
		t.Log(i)
	}
}

func TestForWithMultiVariable(t *testing.T) {

	sum := 0
	for i, j := 0, 0; i < 5 && j < 6; i, j = i+1, j+1 {
		sum += (i + j)
	}
	t.Log(sum)
}

func TestForOnlyRetainCondition(t *testing.T) {

	i := 0
	for i < 10 {
		t.Log(i)
		i++
	}
}

func TestInfiniteFor(t *testing.T) {

	i := 0
	for {
		i = i + 1
		if i > 10 {
			break
		}
		t.Log(i)
	}
}

func TestForRange(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		t.Logf("index is %d, value is %d", i, v)
	}
}

func TestForRangePreferIndex(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5}
	for i := range arr {
		t.Logf("index is %d", i)
	}
}

func TestForRangePreferValue(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5}
	for _, v := range arr {
		t.Logf("value is %d", v)
	}
}

func TestForRangeOnly(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5}
	for range arr {
		t.Logf("range only")
	}
}

func TestForString(t *testing.T) {

	for i, v := range "hello" {
		t.Logf("index is %d, value is %c", i, v)
	}
}

func TestForMap(t *testing.T) {

	aMap := map[string]int{
		"abc": 1,
	}
	aMap["a"] = 1
	aMap["b"] = 2

	for s := range aMap {
		t.Logf("key is %s", s)
	}

	for k, v := range aMap {
		t.Logf("key is %s, value is %d", k, v)
	}
}

func TestForLoopLabel(t *testing.T) {

myLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			t.Logf("i is %d, j is %d", i, j)
			if j == 5 {
				break myLoop
			}
		}
	}
}

func TestSwitch(t *testing.T) {

	a := 10
	switch a {
	case 1:
		t.Log("a is 1")
	case 10:
		t.Log("a is 10")
	default:
		t.Log("a is not 1 or 10")
	}
}

func TestSwitchCompare(t *testing.T) {

	str := "hello"
	switch str {
	case "hello":
		t.Log("str is hello")
	case "world":
		t.Log("str is world")
	default:
		t.Log("str is not hello or world")
	}
}

func TestSwitchWithInitStmt(t *testing.T) {

	switch a := 10; a {
	case 1:
		t.Log("a is 1")
	case 10:
		t.Log("a is 10")
	default:
		t.Log("a is not 1 or 10")
	}
}

func TestSwitchCaseMultiValues(t *testing.T) {

	switch a := 100; a {
	case 1, 10, 100:
		t.Log("a is 1 or 10 or 100")
	case 2:
		t.Log("a is 2")
	default:
		t.Log("a is not 1 or 10 or 100")
	}
}

func TestSwitchRunNextCase(t *testing.T) {

	switch a := 1; a {
	case 1:
		t.Log("a is 1")
		fallthrough
	case 2:
		t.Log("a is 2")
	default:
		t.Log("a is not 1 or 2")
	}
}

func TestSwitchType(t *testing.T) {

	var x interface{}
	x = "hello"
	switch x.(type) {
	case string:
		t.Log("x is string")
	case int:
		t.Log("x is int")
	default:
		t.Log("x is not string or int")
	}
}

func TestSwitchTypeValue(t *testing.T) {

	var x interface{}
	x = "hello"
	switch v := x.(type) {
	case string:
		t.Log("x is string, value is ", v)
	case int:
		t.Log("x is int, value is ", v)
	default:
		t.Log("x is not string or int")
	}
}
