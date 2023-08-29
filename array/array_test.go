package array_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayDeclare(t *testing.T) {
	var arr [5]int
	t.Log(arr)
}

func TestArrayInit(t *testing.T) {
	var emptyArr [5]int
	t.Log(emptyArr)

	arr := [5]int{1, 2, 3, 4, 5}
	t.Log(arr)

	arr1 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr1)
}

func TestArrayEquals(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	arr1 := [...]int{1, 2, 3, 4, 5}
	assert.Equal(t, arr, arr1, "The two words should be the same.")

	doCalculate(arr)
	doCalculate(arr1)

	//arr2 := [6]int{1, 2, 3, 4, 5, 65}
	//do_calculate(arr2)  //长度不同，编译不通过

	//arr3 := [5]int32{1, 2, 3, 4, 5}
	//do_calculate(arr3)  //类型不同，编译不通过
}

func doCalculate(arr [5]int) {
	arr[0] = 100
}

func TestArrayModifyEle(t *testing.T) {
	arr := [5]int{
		2: 101,
	}
	t.Log(arr)
}

func TestArraySlice(t *testing.T) {
	var s0 []int
	t.Log(s0, len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0))

	arr := []int{1, 2, 3, 4, 5}
	t.Log(arr, len(arr), cap(arr))

	s2 := make([]int, 3, 5)
	t.Log(s2, len(s2), cap(s2))
	s2 = append(s2, 1)
	t.Log(s2, len(s2), cap(s2))
}

func TestSliceFromArray(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	t.Log(slice)
}

func TestSliceFromArray2(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]
	t.Log(slice)
}

func TestSliceAppend(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	slice = append(slice, 6)
	t.Log(slice)
	t.Log(arr)
}

func TestDeclareSliceByNew(t *testing.T) {
	var slice = make([]int, 2, 4)
	t.Log(slice)
}

func TestSliceDynamicExpend(t *testing.T) {
	var s []int
	s = append(s, 11)
	t.Log(len(s), cap(s))
	s = append(s, 12)
	t.Log(len(s), cap(s))
	s = append(s, 13)
	t.Log(len(s), cap(s))
	s = append(s, 14)
	t.Log(len(s), cap(s))
	assert.Equal(t, 4, cap(s), "The two words should be the same.")

	s = append(s, 15)
	t.Log(len(s), cap(s))
	assert.Equal(t, 8, cap(s), "The two words should be the same.")
}

func TestSliceDataUnbind(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	t.Log(arr)
	slice := arr[1:3]
	t.Log("len: ", len(slice), "cap: ", cap(slice), "data: ", slice)
	slice = append(slice, 16)
	t.Log("after append 16 len: ", len(slice), "cap: ", cap(slice), "data: ", slice)
	slice = append(slice, 17)
	t.Log("after append 17 len: ", len(slice), "cap: ", cap(slice), "data: ", slice)
	slice = append(slice, 18)
	t.Log("after append 18 len: ", len(slice), "cap: ", cap(slice), "data: ", slice)

	slice[0] = 202
	t.Log("after modify slice[0] len: ", len(slice), "cap: ", cap(slice), "data: ", slice, "old arr: ", arr)

}

func TestSliceToArrayPointer(t *testing.T) {
	slice := []int32{1, 2, 3}
	// from go 1.17, convert array length must be less then or equals slice length
	arr := (*[3]int32)(slice)
	t.Log("arr", arr)
	arr[0] = 4
	t.Log("arr", arr)
	t.Log(arr[0])
}

func TestSliceToArrayPointerLessThanSliceLength(t *testing.T) {
	slice := []int32{1, 2, 3}
	// from go 1.17, convert array length must be less then or equals slice length
	arr := (*[2]int32)(slice)
	t.Log("arr", arr)
	arr[0] = 4
	t.Log("arr", arr)
	t.Log(arr[0])
}

func TestNilSliceToArrayPointer(t *testing.T) {
	var slice []int32
	arr := (*[0]int32)(slice)
	// arr is nil
	t.Log(arr)
	assert.Nil(t, arr)
}

func TestEmptySliceToArrayPointer(t *testing.T) {
	slice := []int32{}
	arr := (*[0]int32)(slice)
	t.Log(arr)
	assert.Equal(t, 0, len(arr))
}

func TestArrayTravel(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	t.Log("Travel by range")

	for idx, ele := range arr {
		t.Log(idx, ele)
	}
}

func TestSliceGrowing(t *testing.T) {
	var slice []int
	for i := 0; i < 10; i++ {
		slice = append(slice, i)
		t.Log(len(slice), cap(slice))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	assert.Equal(t, a, b, "The two words should be the same.")
}
