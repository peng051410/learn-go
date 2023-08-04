package trace_test

import (
	"github.com/peng051410/instrument_trace"
)

func a() {
	defer trace.Trace()()
	b()
}

func b() {
	defer trace.Trace()()
	c()
}
func c() {
	defer trace.Trace()()
	d()
}
func d() {
	defer trace.Trace()()
}

func ExampleTrace() {
	a()
	// Output:
	//g[00001]:  -->github.com/peng051410/instrument_trace_test.a
	//g[00001]:   -->github.com/peng051410/instrument_trace_test.b
	//g[00001]:    -->github.com/peng051410/instrument_trace_test.c
	//g[00001]:     -->github.com/peng051410/instrument_trace_test.d
	//g[00001]:     <--github.com/peng051410/instrument_trace_test.d
	//g[00001]:    <--github.com/peng051410/instrument_trace_test.c
	//g[00001]:   <--github.com/peng051410/instrument_trace_test.b
	//g[00001]:  <--github.com/peng051410/instrument_trace_test.a
}
