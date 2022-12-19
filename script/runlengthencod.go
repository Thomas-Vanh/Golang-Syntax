package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("type your the string to be encode: \n")
	scanner.Scan()
	input = scanner.Text()
	fmt.Println(runlengthencode(input))

}

func runlengthencode(input string) string {
	var result strings.Builder
	for len(input) > 0 {
		firstletter := input[0]
		inputlength := len(input)
		input = strings.TrimLeft(input, string(firstletter))
		if counter := inputlength - len(input); counter > 1 {
			result.WriteString(strconv.Itoa(counter))
		}
		result.WriteString(string(firstletter))
	}
	return result.String()
}
