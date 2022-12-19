package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func pangram_test(s string) bool {
	s = strings.ToLower(s)      // Make the string lowercase
	done := make(map[rune]bool) // create a map to keep track of which letters we've seen
	for _, r := range s {       // Loop trough each letter in the string
		if r >= 'a' && r <= 'z' && !done[r] { //if the letter is a letter and we haven't seen it before, mark it as seen
			done[r] = true
		}
	}
	return len(done) == 26 // if w've seen every letter, return true
}

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter a string: ") // prompt the user to enter a string
	scanner.Scan()                  // scan the input and put it in scanner
	input = scanner.Text()          // put the text of scanner in input as a string
	if pangram_test(input) {        // check if the input is a pangram and print the result
		fmt.Println("this is a pangram")
	} else {
		fmt.Println("this not a pangram")
	}

}
