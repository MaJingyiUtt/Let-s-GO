package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	listUser()
}

func listUser() {
	file, _ := os.Open("/etc/passwd")
	rd := bufio.NewReader(file)
	var userMap map[string]string = make(map[string]string)
	for {
		line, err := rd.ReadString('\n')
		splitTab := strings.Split(line, ":")
		if len(splitTab) > 3 {
			user := splitTab[0]
			value := splitTab[2]
			userMap[user] = value
			//pour convertir en int : val,err=strconv.Atoi(value)
		}
		if err == io.EOF {
			fmt.Println(line)
			break
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	for key, val := range userMap {
		fmt.Println(key + " " + val)
	}
	file.Close()
}
