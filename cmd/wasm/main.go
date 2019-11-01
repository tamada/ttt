package main

import (
	"fmt"
	"syscall/js"
)

func print(this js.Value, args []js.Value) interface{} {
	fmt.Println(args)
	return nil
}

func registerCallbacks() {
	js.Global().Set("print", js.FuncOf(print))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
