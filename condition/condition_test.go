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
	} else if d := 20; d> b {
		t.Log("d > b")
	}
}

func TestFor(t *testing.T) {

	for i := 0; i < 5 ; i++ {
		t.Log(i)
	}
}

func TestForWithMultiVariable(t *testing.T) {

	sum := 0
	for i, j := 0, 0; i < 5 && j < 6; i,j = i+1, j+1 {
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

}

func TestForMap(t *testing.T) {

}

func TestForLoopLabel(t *testing.T) {
}

func TestSwitch(t *testing.T) {
}

func TestSwitchComparealbe(t *testing.T) {
}

func TestSwitchWithInitStmt(t *testing.T) {
}

func TestSwitchCaseMultiValues(t *testing.T) {
}

func TestSwitchNotRunNextCase(t *testing.T) {
}

func TestSwitchFallthrough(t *testing.T) {

}

func TestSwitchType(t *testing.T) {
}

func TestSwitchTypeValue(t *testing.T) {
}