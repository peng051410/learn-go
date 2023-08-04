package main

import (
	"flag"
	"fmt"
	"github.com/peng051410/instrument_trace/instrumenter"
	"github.com/peng051410/instrument_trace/instrumenter/ast"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	wrote bool
)

func init() {
	flag.BoolVar(&wrote, "w", false, "write result to (source) file instead of stdout")
}

func usage() {
	fmt.Println("instrument [-w] xxx.go")
	flag.PrintDefaults()
}

func main() {
	fmt.Println(os.Args)
	flag.Usage = usage
	flag.Parse() //解决命令行参数

	if len(os.Args) < 2 {
		usage()
		return
	}

	var file string
	if len(os.Args) == 3 {
		file = os.Args[2]
	}

	if len(os.Args) == 2 {
		file = os.Args[1]
	}

	if filepath.Ext(file) != ".go" {
		fmt.Println("file must be a go file")
		return
	}

	var ins instrumenter.Instrumenter //声明接口类型

	ins = ast.New("github.com/peng051410/instrument_trace", "trace", "Trace") //实例化接口
	newSrc, err := ins.Instrument(file)                                       //调用接口方法
	if err != nil {
		panic(err)
	}

	if newSrc == nil {
		fmt.Printf("no trace added fo %s\n", file)
		return
	}

	if !wrote {
		fmt.Println(string(newSrc)) //print newSrc to stdout
		return
	}

	// write new generate go source to file
	if err = ioutil.WriteFile(file, newSrc, 0666); err != nil {
		fmt.Printf("failed to write %s: %s\n", file, err)
		panic(err)
	}
	fmt.Printf("wrote %s\n", file)
}
