package main

import (
	"fmt"
	"os"
)

func main() {
	for _, filename := range os.Args[1:] {
		cat(filename)
	}
}

func cat(path string) {
	file, _ := os.Open(path)
	buffer := make([]byte, 100)
	size := 1
	for size > 0 {
		size, _ = file.Read(buffer)
		fmt.Println(string(buffer[0:size]))
	}
	file.Close()
}

//pour exécuter : go run cat.go myfile.txt
