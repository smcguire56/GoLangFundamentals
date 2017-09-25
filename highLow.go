package main

import "fmt"

func main() {
	list := []int{7, 1, 8, 47, 4, 5}

	for i := 0; i < cap(list); i++ {
		fmt.Println(list[i])
	}

	fmt.Printf("Highest: %d\n", getHighestNumber(list))
	fmt.Printf("Lowest: %d", getLowestNumber(list))

}

func getHighestNumber(list []int) int {
	var highestCurrently = -99999
	for i := 0; i < cap(list); i++ {
		//fmt.Println(list[i])
		if list[i] > highestCurrently {
			highestCurrently = list[i]
		}
	}
	return highestCurrently
}

func getLowestNumber(list []int) int {
	var lowestCurrently = 99999
	for i := 0; i < cap(list); i++ {
		//fmt.Println(list[i])
		if list[i] < lowestCurrently {
			lowestCurrently = list[i]
		}
	}
	return lowestCurrently
}
