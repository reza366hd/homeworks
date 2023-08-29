package sixteen

import "fmt"

func sixteen() {
	number := 125
	printDigits(number)
}

func printDigits(number int) {
	fmt.Println(number % 10)
	if number >= 10 {
		printDigits(number / 10)

	}

}
