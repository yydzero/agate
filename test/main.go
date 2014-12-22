package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type ia interface {
	Fuck() string
}

type IIA struct {
}

func (a IIA) Fuck() string {
	fmt.Println("fuck")
	return "fuck"
}

type ib interface {
	KFC() string
}

type IIB struct {
}

func (b IIB) KFC() string {
	fmt.Println("KFC")
	return "KFC"
}

func equal(a, b interface{}) bool {
	defer func() {
		recover()
	}()
	return a == b
}

func main() {
	cmd := exec.Command("git", "log")
	fmt.Println("path: ", cmd.Path)
	fmt.Println("path: ", cmd.Args)
	fmt.Println("stdin: ", cmd.Stdin)
	fmt.Println("stdout: ", cmd.Stdout)
	//	fmt.Println("env: ", os.Environ())
	fmt.Println("dev/null: ", os.DevNull)

	input := "Input string\nLine 2"
	cmd.Stdin = strings.NewReader(input)

	if _, ok := cmd.Stdin.(*os.File); ok {
		fmt.Println("cmd.Stdin is os.File")
	} else {
		fmt.Println("cmd.Stdin is not file")
	}

	f, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	cmd.Stdin = f
	if _, ok := cmd.Stdin.(*os.File); ok {
		fmt.Println("cmd.Stdin is os.File")
	} else {
		fmt.Println("cmd.Stdin is not file")
	}

	a := IIA{}
	b := IIB{}

	fmt.Println(equal(a, b))

	fmt.Println(runtime.GOOS)

	fmt.Println("now finished")
}
