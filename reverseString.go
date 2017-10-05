// Reverse String in Go lang.
// https://github.com/smcguire56/GoLangFundamentals :: 2017-10-05
// Sean McGuire

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter String: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text + reverseString(text))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
