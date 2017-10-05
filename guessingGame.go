// guessing game in Go.
// https://github.com/smcguire56/GoLangFundamentals :: 2017-10-05
// Sean McGuire
package main

import "time"
import "fmt"
import "math/rand"

func main() {

	var input int
	var actualNum int
	var numbers [100]int
	var i int
	var found int
	// generate a random number
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	actualNum = r1.Intn(100)

	// print the number
	fmt.Println(actualNum)

	// ask the user to guess the number
	for found == 0 {
		i++
		fmt.Printf("Enter guess %d: ", i)
		fmt.Scanf("%d ", &input)
		numbers[i] = input

		// determine if they got the correct answer
		if input == actualNum {
			fmt.Printf("Correct Guess!, after %d guesses", i)
			found = 1
			break
		} else if input < actualNum {
			fmt.Println("Too small")
		} else if input > actualNum {
			fmt.Println("Too big")
		} else {
			fmt.Println("Try again")
		}
	}
}
