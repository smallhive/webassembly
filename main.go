//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
	"time"
)

var (
	inputChan  chan int
	outputChan chan int
)

type str struct {
	Name string
	Age  int
}

func add(this js.Value, args []js.Value) interface{} {
	inputChan <- args[0].Int() + args[1].Int()
	return <-outputChan
}

func add2(this js.Value, args []js.Value) interface{} {
	inputChan <- args[0].Int()
	inputChan <- args[1].Int()
	return <-outputChan
}

func sequence(this js.Value, args []js.Value) interface{} {
	return []interface{}{1, 2, 3, 4, 5}
}

func nmb(this js.Value, args []js.Value) interface{} {
	return time.Now().Unix()
}

func rstr(this js.Value, args []js.Value) interface{} {
	m := make(map[string]interface{})
	m["Name"] = "Alex"
	m["Age"] = 10
	return m
}

func registerCallbacks() {
	js.Global().Set("currentTimestamp", js.FuncOf(nmb))
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("add2", js.FuncOf(add2))
	js.Global().Set("sequence", js.FuncOf(sequence))
	js.Global().Set("rstr", js.FuncOf(rstr))

	// defer nmbFunc.Release()
}

func summer(src <-chan int, dst chan<- int) {
	i := 0
	for v := range src {
		i += v
		dst <- i
	}
}

func main() {
	c := make(chan interface{})
	inputChan = make(chan int, 2)
	outputChan = make(chan int, 2)

	registerCallbacks()
	println("callbacks registered")

	go summer(inputChan, outputChan)

	<-c
}
