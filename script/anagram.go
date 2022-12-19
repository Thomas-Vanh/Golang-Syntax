package main

import (
	"bufio"
	"fmt"
	"os"
)
func main () {
	scanner := bufio.NewScanner(os.Stdin)
	scanner2:= bufio.NewScanner(os.Stdin)
	fmt.Printf("write a word: \n")
	scanner.Scan()
	word := scanner.Text()
	fmt.Printf("put your list of candidates here: \n")
	scanner2.Scan()
	candidates := [] string {scanner2.Text()}
	fmt.Printf(anagram(word,candidates))
}

 func anagram (w string, c string)  string {
	letters := make(map[rune]bool)

	for _, r := range w {
		letters [r] =true
	}
	for rr:= range c {
		
		}
	}
 }