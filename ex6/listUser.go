package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	listUser()
}

func listUser() {
	file, _ := os.Open("/etc/passwd")
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			fmt.Println(line)
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(line)
	}
	file.Close()
}
