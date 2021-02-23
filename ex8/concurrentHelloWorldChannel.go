package main

import (
	"fmt"
)

func main() {
	chi := make(chan bool)
	done := make(chan bool)
	go hello(chi, done)
	world(chi)
	<-done
}

func hello(chi, done chan bool) {
	<-chi
	fmt.Println("hello")
	done <- true
}

func world(chi chan bool) {
	fmt.Println("world")
	chi <- true
}
