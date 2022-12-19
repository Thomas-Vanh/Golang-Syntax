package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("put u anniversary : [format yyyy-mm-dd]")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	thedate, err := time.Parse("2006-01-02", input)
	if err != nil {
		fmt.Printf("not in the good format sry: %s\n", err)
		return
	}
	gigasecond := time.Second * 1e9
	giga_anniversary := thedate.Add(gigasecond)
	fmt.Printf("Your gigasecond anniversary: %s\n", giga_anniversary.Format("2006-01-02"))
}
