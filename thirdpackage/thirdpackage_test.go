package thirdpackage

import (
	"github.com/samber/lo"
	"testing"
)

func TestLoPackage(t *testing.T) {
	t.Log(lo.IsEmpty("123"))

	if lo.IsEmpty("123") || lo.IsEmpty("333") {
		t.Log("yes")
	}
}
