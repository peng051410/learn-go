package map_study

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMap(t *testing.T) {
	m := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(m)
}

func TestCreateEmptyMap(t *testing.T) {
	m := map[string]string{}
	t.Log(m)
}

func TestMapLength(t *testing.T) {
	m := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(len(m))
}

func TestAddValueToMap(t *testing.T) {
	m := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(m)

	m["age"] = "18"
	t.Log(m)
}

func TestMapModifyValue(t *testing.T) {
	m := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(m)

	m["address"] = "shanghai"
	t.Log(m)
}

func TestMapKeyIsExist(t *testing.T) {
	m := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(m)

	v, ok := m["address"]
	t.Log(v, ok)

	v, ok = m["age"]
	t.Log(v, ok)
}

func TestMapValueIsExist(t *testing.T) {
	m := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(m)

	value := "tomyli"
	//value := "tomyli1"
	for k, v := range m {
		if v == value {
			t.Log("value is exist, key is ", k)
		}
	}
}

func TestMapDelete(t *testing.T) {
	m := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(m)

	delete(m, "address")
	t.Log(m)
}

func TestMapIterator(t *testing.T) {
	m := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(m)

	for k, v := range m {
		t.Log(k, v)
	}
}

func TestMapIteratorKey(t *testing.T) {
	m := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(m)

	for k := range m {
		t.Log(k)
	}
}

func TestMapIteratorValue(t *testing.T) {
	m := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(m)

	for _, v := range m {
		t.Log(v)
	}
}

func TestMapIsEquals(t *testing.T) {
	m1 := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(m1)

	m2 := map[string]string{
		"name":    "tomyli",
		"address": "beijing",
	}
	t.Log(m2)

	//t.Log(m1 == m2) // map can't be compared

	// map can't be compared, but map can be compared with nil
	t.Log(m1 == nil)
}

func TestMapNilDeclare(t *testing.T) {

	var m map[string]string
	t.Log(m)
	assert.Nil(t, m, "m is nil")
	assert.Equal(t, 0, len(m), "The two words should be the same.")
	//m["name"] = "tomyli" // panic: assignment to entry in nil map
}

func TestMapValueIsFunc(t *testing.T) {
	// explain that func is the first class citizen in golang
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapAsSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = true
	mySet[3] = true
	t.Log(mySet)

	delete(mySet, 1)
	t.Log(mySet)

	if mySet[1] {
		t.Log("1 is exist")
	} else {
		t.Log("1 is not exist")
	}

	if mySet[2] {
		t.Log("2 is exist")
	} else {
		t.Log("2 is not exist")
	}
}
