package main

import (
	"fmt"
	"time"
)

func main() {
	go hello()
	world()
	time.Sleep(2 * time.Second)
}

func hello() {
	time.Sleep(1 * time.Second)
	fmt.Println("hello")
}

func world() {
	fmt.Println("world")
}
