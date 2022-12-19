package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		name := os.Args[1]
		if name != "" {
			hiname := "Hello" + " " + name + "!"
			fmt.Println(hiname)
		}
	} else {
		hiname := "Hello, World!"
		fmt.Println(hiname)
	}
}
