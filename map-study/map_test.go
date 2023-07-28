package map_study

import "testing"

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
