package main

import (
	"fmt"
	"strings"
)
func main() {
	phrase := "olly olly in come free"
	seen := strings.Split(phrase," ")
	for z := range seen {
		fmt.Println(seen[z])
	}
}
