package thirteen

import (
	"fmt"
)

func thirteen() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var evenNumbers []int
	var oddNumbers []int
	for num := range numbers {
		if numbers[num]%2 != 0 {
			fmt.Println(num, "%", "2")
			evenNumbers = append(evenNumbers, numbers[num])
		}
	}
	fmt.Println("evenNumbers: ", evenNumbers[0:5])

	for num := range numbers {
		if numbers[num]%2 == 0 {
			fmt.Println(num, "%", "2")
			oddNumbers = append(oddNumbers, numbers[num])
		}
	}
	fmt.Println("oddNumbers: ", oddNumbers[0:5])
}
