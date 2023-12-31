package twelve

import (
	"fmt"
)

func twelve() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var evenNumbers []int
	var oddNumbers []int
	for num := range numbers {
		if numbers[num]%2 != 0 {
			fmt.Println(num, "%", "2")
			evenNumbers = append(evenNumbers, numbers[num])
		}
	}
	fmt.Println("evenNumbers: ", evenNumbers)

	for num := range numbers {
		if numbers[num]%2 == 0 {
			fmt.Println(num, "%", "2")
			oddNumbers = append(oddNumbers, numbers[num])
		}
	}
	fmt.Println("oddNumbers: ", oddNumbers)
}
