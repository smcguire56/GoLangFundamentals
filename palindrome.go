// Palindrome in Go lang.
// https://github.com/smcguire56/GoLangFundamentals :: 2017-10-05
// Sean McGuire

package main

import (
    "bufio"
    "fmt"
    "os"
	"strings"
)
func main() {	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter String: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
    text = strings.Replace(text, "\r\n", "", -1)
	fmt.Printf("%t",isPalindrom(text))
}


func isPalindrom(word string) bool {
	var i1 int = 0
	var i2 int = len(word) - 1

	for i2 >= i1 {
		if word[i1] != word[i2] {
			return false
		}
		i1++
		i2--
	}
	return true
}
