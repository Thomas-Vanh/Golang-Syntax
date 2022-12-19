package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("put the phrase here : ")
	scanner.Scan()
	input = scanner.Text()
	fmt.Println(count_word(input))
}

func count_word(phrase string) map[string]int {
	seen := strings.Split(phrase, " ")
	word := make(map[string]int)
	for z := range seen {
		if _, ok := word[seen[z]]; ok {
			word[seen[z]]++
		} else {
			word[seen[z]] = 1
		}

	}
	return word
}
