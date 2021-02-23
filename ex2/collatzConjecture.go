package main

import "fmt"

func main() {
	collatz(23)
}

func collatz(n int) {
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Println(n)
	}
}
