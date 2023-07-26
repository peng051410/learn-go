package main_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStringJoinWithSymbol(t *testing.T) {
	s := "hello"
	s += ", world"
	t.Log(s)
}

func TestStringJoinWithBuilder(t *testing.T) {
	var builder strings.Builder
	builder.WriteString("hello")
	builder.WriteString("world")
	t.Log(builder.String())
}

func TestStringJoin(t *testing.T) {
	str := []string{"hello", "world", "nihao"}
	//底层用的String builder
	result := strings.Join(str, "-")
	t.Log(result)
}

func TestStringLength(t *testing.T) {
	str := "hello"
	t.Log(len(str))
	assert.Equal(t, 5, len(str), "The two words should be the same.")
}
