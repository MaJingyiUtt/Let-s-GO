package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	wall := make([]bool, 1000)
	wall[rand.Intn(1000)] = true
	resulchan := make(chan int)
	for i := 0; i < 1000; i += 100 {
		go seeker(resulchan, wall, i, i+100)
	}
	fmt.Println(<-resulchan)
}

func seeker(resulchan chan int, wall []bool, deb int, fin int) {
	for i := deb; i < fin; i++ {
		if wall[i] {
			resulchan <- i
			break
		}
	}
}
