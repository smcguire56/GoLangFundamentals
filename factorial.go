// A factorial application in Go.
// https://github.com/smcguire56/GoLangFundamentals :: 2017-10-05
// Sean McGuire
package main

import "fmt"

func main() {

	var input, fact, digit, sum int =  0,1,0,0;

	fmt.Println("Please Enter a number:")

	fmt.Scanf("%d", &input)

	fmt.Println("Number Entered:", input)

	for i := 1; i<=input; i++{
		fact = fact * i
	}

	fmt.Println("Number Factorial:", fact)

	for fact > 0 {
		digit = fact % 10
		sum += digit
		fact /=10
	}

	fmt.Println("Number Summed:", sum)
}
