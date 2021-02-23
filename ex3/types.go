package main

import "fmt"

type MyInt int

func even(i int) bool {
	return i%2 == 0
}

func odd(i int) bool {
	return i%2 != 0
}

type MyIntTab []int
type MyTabInt []int

func size(s []int) int {
	return len(s)
}

func MyIntTabSize(s MyIntTab) int {
	return len(s)
}

func main() {
	fmt.Println(even(2))
	fmt.Println(odd(2))
	var myInt MyInt = 1
	fmt.Println(even(int(myInt)))
	listeA := []int{1, 2}
	fmt.Println(size(listeA))
	//fmt.Println(MyIntTabSize(listeA))
	var listeB MyIntTab = MyIntTab{1, 2}
	//fmt.Println(size(listeB))
	fmt.Println(MyIntTabSize(listeB))
}
